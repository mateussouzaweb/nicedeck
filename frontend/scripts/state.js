// State
window.addEventListener('load', async () => {

    /**
     * Load and show available platforms in the software
     */
    async function loadPlatforms() {

        /** @type {Platform[]} */
        const platforms = await requestJson('GET', '/api/platforms')
        const options = platforms.map((platform) => {
            const console = platform.console.toLowerCase().replaceAll(' ', '-')
            return `<label class="checkbox" title="${platform.console}">
                <input type="checkbox" name="platforms[]" value="${platform.name}" />
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

        const destination = $('#state #platforms')
        destination.innerHTML = options.join('')

    }

    on('#state .select-all', 'click', (event) => {
        event.preventDefault()
        const parent = event.target.closest('.group')
        const inputs = $$('input[type="checkbox"]', parent)
        inputs.map((input) => input.checked = "checked")
    })

    on('#state .clear-all', 'click', (event) => {
        event.preventDefault()
        const parent = event.target.closest('.group')
        const inputs = $$('input[type="checkbox"]', parent)
        inputs.map((input) => input.checked = "")
    })

    on('#state form', 'submit', async (event) => {
        event.preventDefault()

        const form = $('#state form')
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
                await requestJson('POST', '/api/sync/state', body)
            })
        } catch (error) {
            window.showError(error)
        } finally {
            button.disabled = false
        }
    })

    try {
        await loadPlatforms()
    } catch (error) {
        window.showError(error)
    }

})