// Install
window.addEventListener('load', async () => {

    /**
     * Load and show available programs in the software
     */
    async function loadPrograms() {

        /** @type {Program[]} */
        const programs = await requestJson('GET', '/api/programs')
        const options = programs.map((program) => {
            return `<label class="checkbox" title="${program.name}">
                <input type="checkbox" name="programs[]" value="${program.id}" />
                <div class="area">
                    <div class="icon">
                        <img loading="lazy" src="${program.iconUrl}" alt="${program.name}" width="96" height="96" />
                    </div>
                    <div class="info">
                        <b>${program.name}</b><br/>
                        <small>${program.description}</small>
                    </div>
                </div>
            </label>`
        })

        const destination = $('#install #programs')
        destination.innerHTML = options.join('')

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
        event.preventDefault();

        const form = $('#install form')
        const button = $('button[type="submit"]', form)

        const data = new FormData(form)
        const body = JSON.stringify({
            programs: data.getAll('programs[]')
        })

        try {
            button.disabled = true
            await window.runAndCaptureConsole(true, async () => {
                await request('POST', '/api/install', body)
                await request('POST', '/api/library/save')
            })
        } finally {
            button.disabled = false
        }
    })

    await loadPrograms()

})