// ROMs
window.addEventListener('load', async () => {

    async function loadPlatforms() {

        /** @type {Platform[]} */
        const platforms = await requestJson('GET', '/api/platforms')
        const options = platforms.map((platform) => {
            return `<label class="checkbox">
                <input type="checkbox" name="platforms[]" value="${platform.name}" checked="checked" />
                <div class="area">
                    <div class="icon">
                        <img loading="lazy" src="${platform.iconUrl}" width="48" height="48" />
                    </div>
                    <div class="info">
                        <b>${platform.name}</b><br/>
                        <small>${platform.console}</small>
                    </div>
                </div>
            </label>`
        })

        const destination = $('#roms #platforms')
        destination.innerHTML = options.join('')

    }

    on('#roms .select-all', 'click', (event) => {
        event.preventDefault()
        const parent = event.target.closest('.group')
        const inputs = $$('input[type="checkbox"]', parent)
        inputs.map((input) => input.checked = "checked")
    })

    on('#roms .clear-all', 'click', (event) => {
        event.preventDefault()
        const parent = event.target.closest('.group')
        const inputs = $$('input[type="checkbox"]', parent)
        inputs.map((input) => input.checked = "")
    })

    on('#roms form', 'submit', async (event) => {
        event.preventDefault();

        const form = $('#roms form')
        const data = new FormData(form)
        const button = $('button[type="submit"]', form)

        try {
            button.disabled = true
            window.watchConsoleOutput()
            await request('POST', '/api/roms', data)
        } finally {
            window.stopConsoleOutput()
            button.disabled = false
        }
    })

    await loadPlatforms()

})