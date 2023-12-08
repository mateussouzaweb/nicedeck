// Modal
window.addEventListener('load', async () => {

    /**
     * Show modal
     * @param {Element} element
     */
    function showModal(element) {

        element.classList.remove('hidden')
        element.classList.add('visible')

        window.setTimeout(() => {
            element.classList.add('in')
        }, 150)

    }

    /**
     * Hide modal
     * @param {Element} element
     */
    function hideModal(element) {

        element.classList.remove('in')

        window.setTimeout(() => {
            element.classList.remove('visible')
            element.classList.add('hidden')
        }, 150)

    }

    document.addEventListener('click', (event) => {

        /** @type {Element} */
        const target = event.target
        const link = target.closest('[data-modal]')
        const close = target.closest('.close')
        const backdrop = target.closest('.backdrop')
        let modal = target.closest('.modal')

        if (link) {
            modal = document.querySelector(
                link.dataset.modal
            )
        }

        if (!modal) {
            return
        }

        if (close || backdrop) {
            event.preventDefault()
            hideModal(modal)
        } else if (link) {
            event.preventDefault()
            showModal(modal)
        }

    })

    window.showModal = showModal
    window.hideModal = hideModal

})