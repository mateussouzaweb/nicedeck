// Install
window.addEventListener('load', () => {

    on('#install form', 'submit', async (event) => {
        event.preventDefault();

        const form = $('#install form')
        const data = new FormData(form)
        await request('POST', '/api/install', data)
    })

})