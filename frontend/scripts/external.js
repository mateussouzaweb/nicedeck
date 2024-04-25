// External
window.addEventListener('load', async () => {

    // Open external links with browser
    on('a', 'click', async (event) => {
        const link = event.target.closest('a')

        if (!link.hostname) {
            return
        }
        if (link.hostname === location.hostname) {
            return
        }

        event.preventDefault()

        if (link.disabled) {
            return
        }

        const body = JSON.stringify({
            link: link.href
        })

        try {
            link.disabled = true
            await window.runAndCaptureConsole(true, async () => {
                await requestJson('POST', '/api/link/open', body)
            })
        } catch (error) {
            window.showError(error)
        } finally {
            link.disabled = false
        }
    })

})