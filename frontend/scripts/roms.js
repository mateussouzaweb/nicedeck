// ROMs
window.addEventListener('load', async () => {

    /**
     * Load and show available platforms in the software
     */
    async function loadPlatforms() {

        /** @type {PlatformsRequestResult} */
        const request = await requestJson('GET', '/api/platforms')
        const platforms = request.data

        const options = platforms.map((platform) => {
            const console = platform.console.toLowerCase().replaceAll(' ', '-')
            return `<label class="checkbox" title="${platform.console}">
                <input type="checkbox" name="platforms[]" value="${platform.name}" checked="checked" />
                <div class="area">
                    <div class="icon">
                        <img loading="lazy" src="/img/platforms/${console}.png" alt="${platform.console}" width="96" height="96" />
                    </div>
                    <div class="info">
                        <b>${platform.name}</b>
                    </div>
                </div>
            </label>`
        })

        // Empty elements for flexbox
        options.push('<div class="fill"></div>')
        options.push('<div class="fill"></div>')
        options.push('<div class="fill"></div>')
        options.push('<div class="fill"></div>')
        options.push('<div class="fill"></div>')
        options.push('<div class="fill"></div>')
        options.push('<div class="fill"></div>')
        options.push('<div class="fill"></div>')
        options.push('<div class="fill"></div>')
        options.push('<div class="fill"></div>')

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
        event.preventDefault()

        const form = $('#roms form')
        const button = $('button[type="submit"]', form)

        if (button.disabled) {
            return
        }

        const data = new FormData(form)
        const body = JSON.stringify({
            platforms: data.getAll('platforms[]'),
            preferences: data.getAll('preferences[]')
        })

        try {
            button.disabled = true
            await window.runAndCaptureConsole(true, async () => {
                await requestJson('POST', '/api/roms', body)
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
        await loadPlatforms()
    } catch (error) {
        window.showError(error)
    }

})