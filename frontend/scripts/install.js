// Install
window.addEventListener('load', async () => {

    async function loadPrograms(){

        /** @type {Program[]} */
        const programs = await requestJson('GET', '/api/programs')
        const options = programs.map((program) => {
            return `<label class="inline">
                <input type="checkbox" name="programs[]" value="${program.id}" checked="checked" />
                <span>${program.name}</span>
            </label>`
        })

        const destination = $('#install #programs')
        destination.innerHTML = options.join('')

    }

    on('#install form', 'submit', async (event) => {
        event.preventDefault();

        const form = $('#install form')
        const data = new FormData(form)
        await request('POST', '/api/install', data)
    })

    await loadPrograms()

})