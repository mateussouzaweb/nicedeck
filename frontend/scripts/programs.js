// Programs
window.addEventListener('load', async () => {

    /** @type {Program[]} */
    let programs = []

    /**
     * Retrieve program by id
     * @param {String} id
     * @returns {Program}
     */
    function getProgram(id) {
        return programs.find((item) => String(item.id) === String(id))
    }

    /**
     * Render programs in HTML
     */
    async function renderPrograms() {

        const html = []
        const append = (category) => {
            html.push(
            `<section class="group">
                <div class="group-info">
                    <h4>${category}</h4>
                    <p class="mass-actions">
                        <span class="select-all">SELECT ALL</span>
                        <span class="separator">/</span>
                        <span class="clear-all">CLEAR</span>
                    </p>
                </div>`)

            programs.map((program) => {
                if (program.category !== category) {
                    return
                }

                const flags = program.flags.map((flag) => {
                    flag = flag.replaceAll('--', '')
                    return `<span class="tag tag-${flag}">${flag}</span>`
                }).join(' ')

                html.push(
                `<label class="checkbox" title="${program.name}">
                    <input type="checkbox" name="programs[]" value="${program.id}" />
                    <div class="area">
                        <div class="icon">
                            <img loading="lazy" src="/img/programs/${program.id}.png" alt="${program.name}" width="96" height="96" />
                        </div>
                        <div class="info">
                            <b>${program.name}</b><br/>
                            <small>${program.description}</small>
                            <div class="tags">${flags}</div>
                        </div>
                        <div class="details" data-details-program="${program.id}" title="Details">
                            <img src="./img/icons/info.svg" alt="Details" width="20" height="20" />
                        </div>
                    </div>
                </label>`)
            })

            html.push(`<div class="fill"></div>`)
            html.push(`<div class="fill"></div>`)
            html.push(`<div class="fill"></div>`)
            html.push(`</section>`)
        }

        const categories = programs.map((program) => {
            return program.category
        }).filter((value, index, array) => {
            return array.indexOf(value) === index
        }).sort()

        categories.map((category) => append(category))

        const destination = $('#programs #list')
        destination.innerHTML = html.join('')

    }

    /**
     * Render program details modal
     * @param {Program} program
     */
    async function renderProgramDetails(program) {

        const modal = $('#modal-details-program')
        const content = $('.content', modal)

        const tags = program.tags.map((tag) => {
            return `<span class="tag">${tag}</span>`
        }).join(' ')

        const flags = program.flags.map((flag) => {
            code = flag.replaceAll('--', '')
            return `<span class="tag tag-${code}">${flag}</span>`
        }).join(' ')

        const folders = program.folders.map((folder) => {
            return `<span class="tag">${folder}</span>`
        }).join(' ')

        const image = `/img/programs/${program.id}.png`
        const html = `
            <div class="columns">
                <div class="icon">
                    <img loading="lazy" src="${image}" alt="${program.name}" width="96" height="96" />
                </div>
                <div class="info">
                    <h4><b>${program.name}</b></h4>
                    <span>Description:</span> ${program.description}<br/>
                    <span>Website:</span> <a target="_blank" href="${program.website}" title="Visit">${program.website}</a>
                </div>
            </div>
            <div class="info">
                ${tags ? `<div class="tags"><span>Tags:</span> ${tags}</div>` : ``}
                ${flags ? `<div class="tags"><span>Flags:</span> ${flags}</div>` : ``}
                ${folders ? `<div class="tags"><span>Folders:</span> ${folders}</div>` : ``}
            </div>
        `

        modal.dataset.program = program.id
        content.innerHTML = html
        window.showModal(modal)

    }

    /**
     * Load and show available programs in the software
     */
    async function loadPrograms() {

        /** @type {ListProgramsResult} */
        const request = await requestJson('GET', '/api/programs')
        programs = request.data
        renderPrograms()

    }

    on('#programs [data-details-program]', 'click', async (event) => {
        event.preventDefault()

        const button = event.target.closest('[data-details-program]')
        if (button.disabled) {
            return
        }

        const program = getProgram(button.dataset.detailsProgram)
        await renderProgramDetails(program)
    })

    on('#programs .select-all', 'click', (event) => {
        event.preventDefault()
        const parent = event.target.closest('.group')
        const inputs = $$('input[type="checkbox"]', parent)
        inputs.map((input) => input.checked = "checked")
    })

    on('#programs .clear-all', 'click', (event) => {
        event.preventDefault()
        const parent = event.target.closest('.group')
        const inputs = $$('input[type="checkbox"]', parent)
        inputs.map((input) => input.checked = "")
    })

    on('#programs form', 'submit', async (event) => {
        event.preventDefault()

        const form = $('#programs form')
        const button = $('button[type="submit"]', form)

        if (button.disabled) {
            return
        }

        const data = new FormData(form)
        const action = data.getAll('action')

        /** @type {InstallProgramsData|RemoveProgramsData} */
        const body = {
            programs: data.getAll('programs[]'),
            preferences: data.getAll('preferences[]')
        }

        await window.runAndCaptureConsole(button, true, async () => {
            try {
                /** @type {LoadLibraryResult} */
                await requestJson('POST', '/api/library/load')
                /** @type {InstallProgramsResult|RemoveProgramsData} */
                await requestJson('POST', `/api/programs/${action}`, JSON.stringify(body))
            } catch (error) {
                window.showError(error)
            }

            try {
                /** @type {SaveLibraryResult} */
                await requestJson('POST', '/api/library/save')
            } catch (error) {
                window.showError(error)
            }
        })

        try {
            await loadPrograms()
            await window.reloadPlatforms()
            await window.reloadState()
            await window.reloadShortcuts()
        } catch (error) {
            window.showError(error)
        }
    })

    try {
        await loadPrograms()
    } catch (error) {
        window.showError(error)
    }

})