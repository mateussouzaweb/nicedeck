// Setup
window.addEventListener('load', () => {

    on('#setup form', 'submit', async (event) => {
        event.preventDefault()

        const form = $('#setup form')
        const button = $('button[type="submit"]', form)

        if (button.disabled) {
            return
        }

        const data = new FormData(form)
        const body = JSON.stringify({
            installOnMicroSD: data.get('install_on_microsd') === 'Y',
            microSDPath: data.get('microsd_path')
        })

        try {
            button.disabled = true
            await window.runAndCaptureConsole(true, async () => {
                await requestJson('POST', '/api/setup', body)
                await requestJson('POST', '/api/library/load')
                await requestJson('POST', '/api/library/save')
            })
        } catch (error) {
            window.showError(error)
        } finally {
            button.disabled = false
        }
    })

})