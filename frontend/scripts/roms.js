// ROMs
window.addEventListener('load', () => {

    on('#roms form', 'submit', async (event) => {
        event.preventDefault();

        const form = $('#roms form')
        const data = new FormData(form)
        await request('POST', '/api/roms', data)
    })

})