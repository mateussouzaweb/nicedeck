// ROMs
window.addEventListener('load', async () => {

    /**
     * Load and show available platforms in the software
     */
    async function loadPlatforms() {

        /** @type {ListPlatformsResult} */
        const request = await requestJson('GET', '/api/platforms')
        const platforms = []

        request.console.map((platform) => {
            platforms.push({
                key: platform.console.toLowerCase().replaceAll(' ', '-'),
                name: platform.console,
                value: platform.name
            })
        })

        const options = platforms.map((platform) => {
            return `<label class="checkbox" title="${platform.name}">
                <input type="checkbox" name="platforms[]" value="${platform.value}" checked="checked" />
                <div class="area">
                    <div class="icon">
                        <img loading="lazy" src="/img/platforms/${platform.key}.png" alt="${platform.name}" width="96" height="96" />
                    </div>
                    <div class="info">
                        <b>${platform.value}</b>
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

        const destination = $('#roms .platforms')
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

        /** @type {ProcessROMsData} */
        const body = {
            platforms: data.getAll('platforms[]'),
            preferences: data.getAll('preferences[]')
        }

        await window.runAndCaptureConsole(button, true, async () => {
            try {
                /** @type {LoadLibraryResult} */
                await requestJson('POST', '/api/library/load')
                /** @type {ProcessROMsResult} */
                await requestJson('POST', '/api/roms', JSON.stringify(body))
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