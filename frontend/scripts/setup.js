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

        /** @type {RunSetupData} */
        const body = {
            useSymlink: data.get('use_symlink') === 'Y',
            storagePath: data.get('storage_path')
        }

        try {
            button.disabled = true
            await window.runAndCaptureConsole(true, async () => {
                /** @type {RunSetupResult} */
                await requestJson('POST', '/api/setup', JSON.stringify(body))
                /** @type {LoadLibraryResult} */
                await requestJson('POST', '/api/library/load')
                /** @type {SaveLibraryResult} */
                await requestJson('POST', '/api/library/save')
            })
        } catch (error) {
            window.showError(error)
        } finally {
            button.disabled = false
        }
    })

})