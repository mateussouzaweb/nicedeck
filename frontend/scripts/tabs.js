// Tabs
window.addEventListener('load', async () => {

    on('ul.tabs li[data-tab]', 'click', async (event) => {
        event.preventDefault()

        $$('section.tab').forEach((item) => {
            item.classList.remove('active')
        })
        $$('ul.tabs li').forEach((item) => {
            item.classList.remove('active')
        })

        const item = event.target.closest('li')
        item.classList.add('active')
    
        const target = item.dataset.tab
        $('section.tab#' + target).classList.add('active')
    })

})