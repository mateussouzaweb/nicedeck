// Setup
window.addEventListener('load', () => {

    on('#setup form', 'submit', async (event) => {
        event.preventDefault();

        const form = $('#setup form')
        const button = $('button[type="submit"]', form)
        
        const data = new FormData(form)
        const body = JSON.stringify({
            installOnMicroSD: data.get('install_on_microsd') === 'Y',
            microSDPath: data.get('microsd_path')
        })

        try {
            button.disabled = true
            await window.runAndCaptureConsole(async () => {
                await request('POST', '/api/setup', body)
                await request('POST', '/api/library/save')
            })
        } finally {
            button.disabled = false
        }
    })

})