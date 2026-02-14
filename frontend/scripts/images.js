// Images
window.addEventListener('load', async () => {

    /**
     * State for image manipulation
     * @type {Object}
     */
    const state = {
        blockElement: null,
        currentValue: null,
        currentImage: null,
        imageType: 'cover',
        imageWidth: 600,
        imageHeight: 900
    }

    /**
     * Image specifications
     * @type {Object}
     */
    const specifications = {
        banner: { label: 'Banner', width: 920, height: 430 },
        cover: { label: 'Cover', width: 600, height: 900 },
        hero: { label: 'Hero', width: 900, height: 300 },
        icon: { label: 'Icon', width: 192, height: 192 },
        logo: { label: 'Logo', width: 300, height: 120 }
    }

    /**
     * Render image selector block
     * @param {HTMLElement} element
     */
    function renderImageSelectors(element) {
        $$('[data-select-image]', element).map((item) => {

            const fieldKey = item.dataset.selectImage
            const currentValue = item.dataset.currentValue
            const currentImage = item.dataset.currentImage
            const imageType = item.dataset.imageType
            const imageLabel = specifications[imageType].label
            const imageWidth = specifications[imageType].width
            const imageHeight = specifications[imageType].height

            const html = `
                <small>${imageLabel}</small>
                <input type="hidden" name="${fieldKey}" value="${currentValue}" />
                <img loading="lazy" src="${currentImage}"
                    alt="${imageLabel}" width="${imageWidth}" height="${imageHeight}"/>
                <span class="select">${currentValue ? 'Change image' : 'Select image'}</span>
                ${currentValue ? '<span class="remove">&times;</span>' : ''}
            `

            item.innerHTML = html
            item.classList.add('image')

            if (!currentValue) {
                item.classList.add('no-image')
            } else {
                item.classList.remove('no-image')
            }

        })
    }

    /**
     * Render modal to select image based on current image state
     * @param {Object} state
     * @param {String} searchTerm
     */
    function renderModal(state, searchTerm) {

        const modal = $('#modal-select-image')
        const content = $('.content', modal)
        const html = `
            <div class="group">
                <label for="search">Search:</label>
                <textarea id="search" name="search" data-search>${searchTerm}</textarea>
            </div>
            <div class="group">
                <div class="options images ${state.imageType}">
                    <p>${searchTerm
                ? 'Loading images, please wait...'
                : 'Please type the text field to search for images.'
            }</p>
                </div>
            </div>
        `

        content.innerHTML = html
        window.showModal(modal)

        if (searchTerm) {
            searchForResults(searchTerm)
        }

    }

    /**
     * Search for image results based on given search term
     * @param {String} searchTerm
     */
    async function searchForResults(searchTerm) {

        const currentValue = state.currentValue
        const currentImage = state.currentImage
        const imageType = state.imageType
        const imageWidth = state.imageWidth
        const imageHeight = state.imageHeight
        const results = []

        if (currentValue && currentImage) {
            results.push({
                value: currentValue,
                image: currentImage
            })
        }

        try {

            const params = new URLSearchParams({
                type: imageType,
                term: searchTerm
            })

            /** @type {ScrapeDataResult} */
            const request = await requestJson('GET', '/api/scrape?' + params.toString())

            new Array(
                ...request.result.coverUrls,
                ...request.result.bannerUrls,
                ...request.result.heroUrls,
                ...request.result.iconUrls,
                ...request.result.logoUrls
            ).map((item) => {
                if (item && currentValue !== item) {
                    results.push({ value: item, image: item })
                }
            })

        } catch (error) {
            console.error(error)
        }

        const modal = $('#modal-select-image')
        const options = $(`.options`, modal)
        const html = []

        if (!results.length) {
            html.push(`<p>No images could be found for this artwork type with the given search term.</p>`)
        }

        results.forEach((item) => {
            const isSelected = currentValue === item.value
            const checked = isSelected ? 'checked="checked"' : ''
            const imageAlt = isSelected ? 'Current Image' : 'Image'

            html.push(`
                <label class="radio">
                    <input type="radio" name="value"
                        data-image="${item.image}"
                        value="${item.value}" ${checked} />
                    <div class="image">
                        <img loading="lazy" src="${item.image}"
                            alt="${imageAlt}"
                            width="${imageWidth}"
                            height="${imageHeight}"/>
                    </div>
                </label>`)
        })

        options.innerHTML = html.join('')

    }

    on('#modal-select-image [data-search]', 'change', async (event) => {

        const modal = $('#modal-select-image')
        const form = $('form', modal)

        const selectedInput = $('input[name="value"]:checked', form)
        state.currentValue = selectedInput ? selectedInput.value : ''
        state.currentImage = selectedInput ? selectedInput.dataset.image : ''

        const searchTerm = event.target.value
        searchForResults(searchTerm)

    })

    on('#modal-select-image form', 'submit', async (event) => {
        event.preventDefault()

        const modal = $('#modal-select-image')
        const form = $('form', modal)
        const button = $('button[type="submit"]', form)

        if (button.disabled) {
            return
        }

        const selectedInput = $('input[name="value"]:checked', form)
        state.currentValue = selectedInput ? selectedInput.value : ''
        state.currentImage = selectedInput ? selectedInput.dataset.image : ''

        const blockElement = state.blockElement
        blockElement.dataset.currentImage = state.currentImage
        blockElement.dataset.currentValue = state.currentValue

        renderImageSelectors(blockElement.parentElement)
        window.hideModal(modal)

    })

    on('form [data-select-image] .remove', 'click', (event) => {
        event.preventDefault()

        const blockElement = event.target.closest('[data-select-image]')
        blockElement.dataset.currentImage = ''
        blockElement.dataset.currentValue = ''

        renderImageSelectors(blockElement.parentElement)

    })

    on('form [data-select-image] .select', 'click', (event) => {
        event.preventDefault()

        // Retrieve initial search term from shortcut name
        const form = event.target.closest('form')
        const data = new FormData(form)
        const searchTerm = data.get('name')

        // Retrieve image details from rendered block
        const blockElement = event.target.closest('[data-select-image]')
        const imageType = blockElement.dataset.imageType
        const currentValue = blockElement.dataset.currentValue
        const currentImage = blockElement.dataset.currentImage

        const imageElement = $('img', blockElement)
        const imageWidth = parseInt(imageElement.getAttribute('width'))
        const imageHeight = parseInt(imageElement.getAttribute('height'))

        // Set modal state and render it
        state.blockElement = blockElement
        state.currentValue = currentValue
        state.currentImage = currentImage
        state.imageType = imageType
        state.imageWidth = imageWidth
        state.imageHeight = imageHeight

        renderModal(state, searchTerm)

    })

    window.renderImageSelectors = renderImageSelectors

})