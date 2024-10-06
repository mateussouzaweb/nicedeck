// Install
window.addEventListener('load', async () => {

    /**
     * Load and show available programs in the software
     */
    async function loadPrograms() {

        /** @type {ProgramsRequestResult} */
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
                if( program.category !== category ){
                    return
                }

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
                        </div>
                    </div>
                </label>`)
            })

            html.push(`</section>`)
        }

        const categories = programs.map((program) => {
            return program.category
        }).filter((value, index, array) => {
            return array.indexOf(value) === index
        }).sort()

        categories.map((category) => append(category))

        const destination = $('#install #programs')
        destination.innerHTML = html.join('')

    }

    on('#install .select-all', 'click', (event) => {
        event.preventDefault()
        const parent = event.target.closest('.group')
        const inputs = $$('input[type="checkbox"]', parent)
        inputs.map((input) => input.checked = "checked")
    })

    on('#install .clear-all', 'click', (event) => {
        event.preventDefault()
        const parent = event.target.closest('.group')
        const inputs = $$('input[type="checkbox"]', parent)
        inputs.map((input) => input.checked = "")
    })

    on('#install form', 'submit', async (event) => {
        event.preventDefault()

        const form = $('#install form')
        const button = $('button[type="submit"]', form)

        if (button.disabled) {
            return
        }

        const data = new FormData(form)
        const body = JSON.stringify({
            programs: data.getAll('programs[]')
        })

        try {
            button.disabled = true
            await window.runAndCaptureConsole(true, async () => {
                await requestJson('POST', '/api/install', body)
                await requestJson('POST', '/api/library/save')
            })
        } catch (error) {
            window.showError(error)
        } finally {
            button.disabled = false
        }

        try {
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