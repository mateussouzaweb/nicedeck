// Shortcuts
window.addEventListener('load', async () => {

    /** @type {Shortcut[]} */
    let shortcuts = []

    /**
     * Retrieve shortcut by id
     * @param {String} id
     * @returns {Shortcut}
     */
    function getShortcut(id) {
        return shortcuts.find((item) => String(item.id) === String(id))
    }

    /**
     * Load and show available platforms in the software
     */
    async function loadPlatforms() {

        /** @type {ListPlatformsResult} */
        const request = await requestJson('GET', '/api/platforms')
        const platforms = request.data

        const options = platforms.map((platform) => {
            return `<label class="radio" title="${platform.console}">
                <input type="checkbox" name="platforms[]" value="${platform.name}" />
                <span>${platform.name}</span>
            </label>`
        })

        const destination = $('#shortcuts .platforms .dropdown')
        destination.innerHTML = options.join('')

    }

    /**
     * Load and show current list of user shortcuts
     */
    async function loadShortcuts() {

        const button = $('#shortcuts #load')

        if (button.disabled) {
            return
        }

        try {
            button.disabled = true

            /** @type {LoadLibraryResult} */
            const library = await requestJson('POST', '/api/library/load')

            /** @type {ListShortcutsResult} */
            const request = await requestJson('GET', '/api/shortcuts')
            shortcuts = request.data

            const platforms = $$('#shortcuts .platforms input:checked')
            const search = $('#shortcuts .search input')
            const filters = platforms.map((input) => {
                return input.value
            })

            const items = shortcuts.filter((shortcut) => {
                if (!filters.length) {
                    return true
                }
                for (const filter of filters) {
                    if (shortcut.name.includes("[" + filter + "]")) {
                        return true
                    }
                }
                return false
            })
            .filter((shortcut) => {
                if (!search.value.length) {
                    return true
                }
                return String(shortcut.name).toLowerCase().includes(
                    String(search.value).toLowerCase()
                )
            })
            .filter((shortcut) => {
                return shortcut !== null
            })
            .map((shortcut) => {
                const coverUrl = (shortcut.coverPath)
                    ? String(shortcut.coverPath).replace(library.imagesPath, "/grid/image")
                    : './img/default/cover.png'

                return `<article class="item shortcut" title="${shortcut.name}">
                    <div class="area">
                        <div class="image">
                            <img loading="lazy" src="${coverUrl}" alt="${shortcut.name}" width="600" height="900"/>
                        </div>
                        <div class="info">
                            <div class="title">
                                <h4>${shortcut.name}</h4>
                            </div>
                            <div class="actions">
                                <button type="button" data-launch-shortcut="${shortcut.id}" title="Launch">
                                    <img src="./img/icons/launch.svg" alt="Launch" width="24" height="24" />
                                </button>
                                <button type="button" data-update-shortcut="${shortcut.id}" title="Update">
                                    <img src="./img/icons/update.svg" alt="Update" width="24" height="24" />
                                </button>
                                <button type="button" data-delete-shortcut="${shortcut.id}" title="Delete">
                                    <img src="./img/icons/delete.svg" alt="Delete" width="24" height="24" />
                                </button>
                            </div>
                        </div>
                    </div>
                </article>`
            })

            if (!items.length) {
                items.push(`<article class="item message">
                    <div class="area">
                        No library shortcuts to display here yet...
                    </div>
                </article>`)
            } else {
                // Empty elements for flexbox
                items.push('<div class="fill"></div>')
                items.push('<div class="fill"></div>')
                items.push('<div class="fill"></div>')
                items.push('<div class="fill"></div>')
                items.push('<div class="fill"></div>')
            }

            const destination = $('#shortcuts #list')
            destination.innerHTML = items.join('')
        } finally {
            button.disabled = false
        }

    }

    on('#shortcuts #filters input', 'change', async () => {
        try {
            await loadShortcuts()
        } catch (error) {
            window.showError(error)
        }
    })

    on('#shortcuts > form', 'submit', async (event) => {
        event.preventDefault()

        try {
            await loadShortcuts()
        } catch (error) {
            window.showError(error)
        }
    })

    on('#shortcuts [data-launch-shortcut]', 'click', async (event) => {
        event.preventDefault()

        const button = event.target.closest('[data-launch-shortcut]')

        if (button.disabled) {
            return
        }

        const modal = $('#modal-launch-shortcut')
        const content = $('.content', modal)
        const shortcut = getShortcut(button.dataset.launchShortcut)

        modal.dataset.shortcut = shortcut.id
        content.innerHTML = `<p>Launching <b>${shortcut.name}</b>...</p>`
        window.showModal(modal)

        /** @type {LaunchShortcutData} */
        const body = {
            id: shortcut.id
        }

        await window.runAndCaptureConsole(button, false, async () => {
            try {
                /** @type {LaunchShortcutResult} */
                await requestJson('POST', '/api/shortcut/launch', JSON.stringify(body))
            } catch (error) {
                window.showError(error)
            }
        })

        window.setTimeout(() => {
            window.hideModal(modal)
        }, 1000)
    })

    on('#shortcuts [data-update-shortcut]', 'click', async (event) => {

        event.preventDefault()
        const element = event.target.closest('[data-update-shortcut]')
        const shortcut = getShortcut(element.dataset.updateShortcut)

        const modal = $('#modal-update-shortcut')
        const content = $('.content', modal)
        const html = `
            <div class="group">
                <label for="program">Program:</label>
                <textarea id="program" name="program">${shortcut.program}</textarea>
            </div>
            <div class="group">
                <label for="name">Name:</label>
                <textarea id="name" name="name">${shortcut.name}</textarea>
            </div>
            <div class="group">
                <label for="description">Description:</label>
                <textarea id="description" name="description">${shortcut.description}</textarea>
            </div>
            <div class="group">
                <label for="startDirectory">Start Directory:</label>
                <textarea id="startDirectory" name="startDirectory">${shortcut.startDirectory}</textarea>
            </div>
            <div class="group">
                <label for="executable">Executable:</label>
                <textarea id="executable" name="executable">${shortcut.executable}</textarea>
            </div>
            <div class="group">
                <label for="launchOptions">Launch Options:</label>
                <textarea id="launchOptions" name="launchOptions">${shortcut.launchOptions}</textarea>
            </div>
            <div class="group">
                <label for="tags">Tags:</label>
                <textarea id="tags" name="tags">${shortcut.tags.join(',')}</textarea>
            </div>
            <div class="images">
                <p>Loading images, please wait...</p>
            </div>
        `

        modal.dataset.shortcut = shortcut.id
        content.innerHTML = html
        window.showModal(modal)

        const nameInput = $('#name', content)
        const changeEvent = new CustomEvent('change', {
            bubbles: true
        })

        nameInput.dispatchEvent(changeEvent)

    })

    on('#shortcuts [data-delete-shortcut]', 'click', (event) => {

        event.preventDefault()
        const element = event.target.closest('[data-delete-shortcut]')
        const shortcut = getShortcut(element.dataset.deleteShortcut)

        const modal = $('#modal-delete-shortcut')
        const content = $('.content', modal)

        modal.dataset.shortcut = shortcut.id
        content.innerHTML = `<p>Are you sure you want to delete the shortcut to <b>${shortcut.name}</b>?</p>`
        window.showModal(modal)

    })

    on('#shortcuts #modal-update-shortcut #name', 'change', async (event) => {

        const modal = $('#modal-update-shortcut')
        const subContent = $('.content .images', modal)

        const shortcut = getShortcut(modal.dataset.shortcut)
        const term = event.target.value || shortcut.name

        try {

            /** @type {ScrapeDataResult} */
            const request = await requestJson('GET', '/api/scrape?term=' + encodeURIComponent(term))
            const scrape = request.result

            const html = []
            const append = (type, title, selected, images, width, height) => {
                html.push(
                `<section class="group group-${type}">
                    <h4>${title}:</h4>`)

                if (!images || !images.length) {
                    html.push(`<p class="alert">No images were found for this artwork type.</p>`)
                } else {
                    html.push(`<div class="options">`)
                    images.forEach((item, index) => {
                        const checked = selected === item ? 'checked="checked"' : ''
                        html.push(
                        `<label class="radio">
                            <input type="radio" name="${type}" value="${item}" ${checked} />
                            <div class="image">
                                <img loading="lazy" src="${item}" alt="Image ${index}"
                                width="${width}" height="${height}"/>
                            </div>
                        </label>`)
                    })
                    html.push(
                        `<label class="radio">
                            <input type="radio" name="${type}" value="" ${!selected ? 'checked="checked"' : ''} />
                            <div class="image">
                                <div class="no-image">No Image</div>
                            </div>
                        </label>
                    </div>`)
                }

                html.push(`</section>`)
            }

            append('cover', 'Cover Artworks', shortcut.coverUrl, scrape.coverUrls, 600, 900)
            append('banner', 'Banner Artworks', shortcut.bannerUrl, scrape.bannerUrls, 920, 430)
            append('hero', 'Hero Artworks', shortcut.heroUrl, scrape.heroUrls, 600, 900)
            append('icon', 'Icon Artworks', shortcut.iconUrl, scrape.iconUrls, 192, 192)
            append('logo', 'Logo Artworks', shortcut.logoUrl, scrape.logoUrls, 600, 900)

            subContent.innerHTML = html.join('')

        } catch (error) {
            window.showError(error)
            window.hideModal(modal)
        }

    })

    on('#shortcuts #modal-update-shortcut form', 'submit', async (event) => {
        event.preventDefault()

        const modal = $('#modal-update-shortcut')
        const form = $('form', modal)
        const button = $('button[type="submit"]', form)

        if (button.disabled) {
            return
        }

        const shortcut = getShortcut(modal.dataset.shortcut)
        const data = new FormData(form)

        /** @type {ModifyShortcutData} */
        const body = {
            action: 'update',
            id: shortcut.id,
            program: data.get('program'),
            name: data.get('name'),
            description: data.get('description'),
            startDirectory: data.get('startDirectory'),
            executable: data.get('executable'),
            launchOptions: data.get('launchOptions'),
            iconUrl: data.get('icon'),
            logoUrl: data.get('logo'),
            coverUrl: data.get('cover'),
            bannerUrl: data.get('banner'),
            heroUrl: data.get('hero'),
            tags: data.get('tags').split(',')
        }

        await window.runAndCaptureConsole(button, false, async () => {
            try {
                /** @type {ModifyShortcutResult} */
                await requestJson('POST', '/api/shortcut/modify', JSON.stringify(body))
                /** @type {SaveLibraryResult} */
                await requestJson('POST', '/api/library/save')
                await loadShortcuts()
            } catch (error) {
                window.showError(error)
            }
        })

        window.hideModal(modal)

    })

    on('#shortcuts #modal-delete-shortcut form', 'submit', async (event) => {
        event.preventDefault()

        const modal = $('#modal-delete-shortcut')
        const form = $('form', modal)
        const button = $('button[type="submit"]', form)

        if (button.disabled) {
            return
        }

        const shortcut = getShortcut(modal.dataset.shortcut)

        /** @type {ModifyShortcutData} */
        const body = {
            action: 'delete',
            id: shortcut.id
        }

        await window.runAndCaptureConsole(button, false, async () => {
            try {
                /** @type {ModifyShortcutResult} */
                await requestJson('POST', '/api/shortcut/modify', JSON.stringify(body))
                /** @type {SaveLibraryResult} */
                await requestJson('POST', '/api/library/save')
                await loadShortcuts()
            } catch (error) {
                window.showError(error)
            }
        })

        window.hideModal(modal)

    })

    on('#shortcuts #load', 'click', async (event) => {
        event.preventDefault()

        try {
            await loadShortcuts()
        } catch (error) {
            window.showError(error)
        }
    })

    try {
        await loadShortcuts()
    } catch (error) {
        window.showError(error)
    }

    try {
        await loadPlatforms()
    } catch (error) {
        window.showError(error)
    }

    window.loadShortcuts = loadShortcuts

})