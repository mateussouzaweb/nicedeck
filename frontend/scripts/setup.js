// Setup
window.addEventListener('load', () => {

    on('#setup form', 'submit', async (event) => {
        event.preventDefault();

        const form = $('#setup form')
        const data = new FormData(form)
        await request('POST', '/api/setup', data)
    })

})