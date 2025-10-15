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

        /** @type {OpenLinkData} */
        const body = {
            link: link.href
        }

        await window.runAndCaptureConsole(link, false, async () => {
            try {
                /** @type {OpenLinkResult} */
                await requestJson('POST', '/api/link/open', JSON.stringify(body))
            } catch (error) {
                window.showError(error)
            }
        })
    })

})