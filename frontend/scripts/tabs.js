// Tabs
window.addEventListener('load', async () => {

    on('ul.tabs li', 'click', async (event) => {
        event.preventDefault()

        $$('section.tab').forEach((item) => {
            item.classList.remove('active')
        })

        const target = event.target.closest('li').dataset.tab
        $('section.tab#' + target).classList.add('active')
    })

})