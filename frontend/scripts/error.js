// Error output
window.addEventListener('load', async () => {

    /**
     * Show error message
     * @param {Error} error 
     */
    function showError(error) {

        const area = $('#error')
        const html = `<div class="alert">
            <p>${error.message}</p>
            <span id="close">&times;</span>
        </div>`

        area.innerHTML = html
        area.scrollIntoView({
            behavior: 'smooth',
            block: 'start',
            inline: 'nearest'
        })

    }

    on('#error #close', 'click', async (event) => {
        event.preventDefault()
        $('#error').innerHTML = ''
    })

    window.showError = showError

})