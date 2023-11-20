// ROMs
window.addEventListener('load', async () => {

    async function loadPlatforms() {

        /** @type {Platform[]} */
        const platforms = await requestJson('GET', '/api/platforms')
        const options = platforms.map((platform) => {
            return `<label class="inline">
                <input type="checkbox" name="platforms[]" value="${platform.name}" checked="checked" />
                <span>${platform.console}</span>
            </label>`
        })

        const destination = $('#roms #platforms')
        destination.innerHTML = options.join('')

    }

    on('#roms form', 'submit', async (event) => {
        event.preventDefault();

        const form = $('#roms form')
        const data = new FormData(form)
        await request('POST', '/api/roms', data)
    })

    await loadPlatforms()

})