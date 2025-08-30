// Programs
window.addEventListener('load', async () => {

    /**
     * Load and show available programs in the software
     */
    async function loadPrograms() {

        /** @type {ListProgramsResult} */
        const request = await requestJson('GET', '/api/programs')
        const programs = request.data

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

                var classes = program.flags.join(' ').replaceAll('--', '');
                html.push(
                `<label class="checkbox ${classes}" title="${program.name}">
                    <input type="checkbox" name="programs[]" value="${program.id}" />
                    <div class="area">
                        <div class="icon">
                            <img loading="lazy" src="/img/programs/${program.id}.png" alt="${program.name}" width="96" height="96" />
                        </div>
                        <div class="info">
                            <b>${program.name}</b><br/>
                            <small>${program.description}</small>
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
            await window.loadShortcuts()
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