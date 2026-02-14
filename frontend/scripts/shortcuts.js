// Shortcuts
window.addEventListener('load', async () => {

    /** @type {Library[]} */
    let library = {}

    /** @type {ConsolePlatform[]} */
    let consolePlatforms = []

    /** @type {NativePlatform[]} */
    let nativePlatforms = []

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
     * Retrieve image with given path
     * @param {String} type
     * @param {String} path
     * @returns {String}
     */
    function getImage(type, path) {
        path = (path) ? String(path) : `/img/default/${type}.png`
        const url = path.replace(library.imagesPath, `/grid/image`)
        const domain = window.location.origin
        return domain + url + '?t=' + library.timestamp
    }

    /**
     * Check if text contains given term
     * @param {String} text
     * @param {String} term
     * @returns {Boolean}
     */
    function containsText(text, term) {
        return String(text).toLowerCase().includes(
            String(term).toLowerCase()
        )
    }

    /**
     * Render filters in HTML
     */
    async function renderFilters() {

        const inputs = $$('#shortcuts .tags input:checked')
        const active = inputs.map((input) => {
            return input.value
        })

        const tags = []
        consolePlatforms.map((platform) => {
            tags.push(platform.name)
        })
        nativePlatforms.map((platform) => {
            tags.push(platform.runtime)
        })
        shortcuts.map((shortcut) => {
            tags.push(...shortcut.tags)
        })

        const unique = [...new Set(tags)].filter(Boolean).sort()
        const options = unique.map((tag) => {
            const checked = active.includes(tag) ? 'checked="checked"' : ''
            return `<label class="radio" title="${tag}">
                <input type="checkbox" name="tags[]"
                    value="${tag}" ${checked} />
                <span>${tag}</span>
            </label>`
        })

        const destination = $('#shortcuts .tags .dropdown')
        destination.innerHTML = options.join('')

    }

    /**
     * Render shortcuts in HTML
     */
    async function renderShortcuts() {

        const form = $('#shortcuts #filters')
        const data = new FormData(form)
        const filters = {
            search: data.get('search'),
            tags: data.getAll('tags[]')
        }

        const filtered = shortcuts.filter((shortcut) => {
            let valid = true

            if (filters.tags.length) {
                valid = valid && filters.tags.every((tag) => {
                    return shortcut.tags.includes(tag)
                })
            }
            if (filters.search) {
                valid = valid && containsText(
                    shortcut.name, filters.search
                )
            }

            return valid
        }).filter((shortcut) => {
            return shortcut !== null
        })

        const items = filtered.map((shortcut) => {
            const coverImage = getImage('cover', shortcut.coverPath)

            return `
            <article class="item shortcut" title="${shortcut.name}">
                <div class="area">
                    <div class="image">
                        <img loading="lazy" src="${coverImage}" alt="${shortcut.name}" width="600" height="900"/>
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

    }

    /**
     * Render shortcut create modal
     */
    async function renderCreateShortcut() {

        const modal = $('#modal-create-shortcut')
        const content = $('.content', modal)
        const html = `
            <div class="group">
                <label for="name">Name:</label>
                <textarea id="name" name="name"></textarea>
            </div>
            <div class="group">
                <label for="path">Path:</label>
                <textarea id="path" name="path"></textarea>
            </div>`

        content.innerHTML = html
        window.showModal(modal)

    }

    /**
     * Render shortcut add modal
     */
    async function renderAddShortcut() {

        const modal = $('#modal-add-shortcut')
        const content = $('.content', modal)
        const html = `
            <div class="group">
                <label for="id">ID:</label>
                <textarea id="id" name="id"></textarea>
            </div>
            <div class="group">
                <label for="name">Name:</label>
                <textarea id="name" name="name" required data-search-artworks></textarea>
            </div>
            <div class="group">
                <h4>Images:</h4>
                <div class="images">
                    <div class="column">
                        <div data-select-image="cover"
                            data-current-value=""
                            data-current-image=""></div>
                    </div>
                    <div class="column">
                        <div data-select-image="banner"
                            data-current-value=""
                            data-current-image=""></div>
                        <div data-select-image="hero"
                            data-current-value=""
                            data-current-image=""></div>
                        <div data-select-image="logo"
                            data-current-value=""
                            data-current-image=""></div>
                    </div>
                    <div class="column icon">
                        <div data-select-image="icon"
                            data-current-value=""
                            data-current-image=""></div>
                    </div>
                </div>
            </div>
            <div class="group">
                <label for="description">Description:</label>
                <textarea id="description" name="description"></textarea>
            </div>
            <div class="group">
                <label for="tags">Tags:</label>
                <textarea id="tags" name="tags"></textarea>
            </div>
            <div class="group">
                <label for="program">Program:</label>
                <textarea id="program" name="program"></textarea>
            </div>
            <div class="group">
                <label for="startDirectory">Start Directory:</label>
                <textarea class="resizable" id="startDirectory" name="startDirectory" required></textarea>
            </div>
            <div class="group">
                <label for="executable">Executable:</label>
                <textarea class="resizable" id="executable" name="executable" required></textarea>
            </div>
            <div class="group">
                <label for="launchOptions">Launch Options:</label>
                <textarea class="resizable" id="launchOptions" name="launchOptions"></textarea>
            </div>
            <div class="group">
                <label for="relativePath">Relative Path:</label>
                <textarea class="resizable" id="relativePath" name="relativePath"></textarea>
            </div>            
        `

        content.innerHTML = html
        window.showModal(modal)
        window.renderImageSelectors(modal)

    }

    /**
     * Render shortcut update modal
     * @param {Shortcut} shortcut
     */
    async function renderUpdateShortcut(shortcut) {

        const modal = $('#modal-update-shortcut')
        const subTitle = $('.header h3 small', modal)
        const content = $('.content', modal)

        const coverImage = getImage('cover', shortcut.coverPath)
        const bannerImage = getImage('banner', shortcut.bannerPath)
        const heroImage = getImage('hero', shortcut.heroPath)
        const iconImage = getImage('icon', shortcut.iconPath)
        const logoImage = getImage('logo', shortcut.logoPath)
        const html = `
            <div class="group">
                <label for="name">Name:</label>
                <textarea id="name" name="name" required data-search-artworks>${shortcut.name}</textarea>
            </div>
            <div class="group">
                <span>Images:</span>
                <div class="images">
                    <div class="column">
                        <div data-select-image="cover"
                            data-current-value="${shortcut.coverPath}"
                            data-current-image="${coverImage}"></div>
                    </div>
                    <div class="column">
                        <div data-select-image="banner"
                            data-current-value="${shortcut.bannerPath}"
                            data-current-image="${bannerImage}"></div>
                        <div data-select-image="hero"
                            data-current-value="${shortcut.heroPath}"
                            data-current-image="${heroImage}"></div>
                        <div data-select-image="logo"
                            data-current-value="${shortcut.logoPath}"
                            data-current-image="${logoImage}"></div>
                    </div>
                    <div class="column icon">
                        <div data-select-image="icon"
                            data-current-value="${shortcut.iconPath}"
                            data-current-image="${iconImage}"></div>
                    </div>
                </div>
            </div>
            <div class="group">
                <label for="description">Description:</label>
                <textarea id="description" name="description">${shortcut.description}</textarea>
            </div>
            <div class="group">
                <label for="tags">Tags:</label>
                <textarea id="tags" name="tags">${shortcut.tags.join(',')}</textarea>
            </div>
            <div class="group">
                <label for="program">Program:</label>
                <textarea id="program" name="program">${shortcut.program}</textarea>
            </div>
            <div class="group">
                <label for="startDirectory">Start Directory:</label>
                <textarea class="resizable" id="startDirectory" name="startDirectory" required>${shortcut.startDirectory}</textarea>
            </div>
            <div class="group">
                <label for="executable">Executable:</label>
                <textarea class="resizable" id="executable" name="executable" required>${shortcut.executable}</textarea>
            </div>
            <div class="group">
                <label for="launchOptions">Launch Options:</label>
                <textarea class="resizable" id="launchOptions" name="launchOptions">${shortcut.launchOptions}</textarea>
            </div>
            <div class="group">
                <label for="relativePath">Relative Path:</label>
                <textarea class="resizable" id="relativePath" name="relativePath">${shortcut.relativePath}</textarea>
            </div>
        `

        modal.dataset.shortcut = shortcut.id
        subTitle.innerText = `(${shortcut.id})`
        content.innerHTML = html

        window.showModal(modal)
        window.renderImageSelectors(modal)

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
            const libraryRequest = await requestJson('POST', '/api/library/load')

            /** @type {LoadLibraryResult} */
            const platformsRequest = await requestJson('GET', '/api/platforms')

            /** @type {ListShortcutsResult} */
            const shortcutsRequest = await requestJson('GET', '/api/shortcuts')

            library = libraryRequest.data
            consolePlatforms = platformsRequest.console
            nativePlatforms = platformsRequest.native
            shortcuts = shortcutsRequest.data

            await renderFilters()
            await renderShortcuts()

        } finally {
            button.disabled = false
        }

    }

    on('#shortcuts #filters input', 'change', async () => {
        try {
            await renderShortcuts()
        } catch (error) {
            window.showError(error)
        }
    })

    on('#shortcuts > form', 'submit', async (event) => {
        event.preventDefault()

        try {
            await renderShortcuts()
        } catch (error) {
            window.showError(error)
        }
    })

    on('#shortcuts #create', 'click', async (event) => {
        event.preventDefault()
        await renderCreateShortcut()
    })

    on('#shortcuts #add', 'click', async (event) => {
        event.preventDefault()
        await renderAddShortcut()
    })

    on('#shortcuts #sync', 'click', async (event) => {
        event.preventDefault()
        const modal = $('#modal-sync-shortcuts')
        window.showModal(modal)
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

        await renderUpdateShortcut(shortcut)

    })

    on('#shortcuts [data-delete-shortcut]', 'click', (event) => {

        event.preventDefault()
        const element = event.target.closest('[data-delete-shortcut]')
        const shortcut = getShortcut(element.dataset.deleteShortcut)

        const modal = $('#modal-delete-shortcut')
        const subTitle = $('.header h3 small', modal)
        const content = $('.content', modal)

        modal.dataset.shortcut = shortcut.id
        subTitle.innerText = `(${shortcut.id})`
        content.innerHTML = `<p>Are you sure you want to delete the shortcut to <b>${shortcut.name}</b>?</p>`
        window.showModal(modal)

    })

    on('#modal-create-shortcut form', 'submit', async (event) => {
        event.preventDefault()

        const modal = $('#modal-create-shortcut')
        const form = $('form', modal)
        const button = $('button[type="submit"]', form)

        if (button.disabled) {
            return
        }

        const data = new FormData(form)

        /** @type {CreateShortcutData} */
        const body = {
            name: data.get('name'),
            path: data.get('path')
        }

        await window.runAndCaptureConsole(button, false, async () => {
            try {
                /** @type {AddShortcutResult} */
                await requestJson('POST', '/api/shortcut/create', JSON.stringify(body))
                /** @type {SaveLibraryResult} */
                await requestJson('POST', '/api/library/save')
                await loadShortcuts()
            } catch (error) {
                window.showError(error)
            }
        })

        window.hideModal(modal)

    })

    on('#modal-add-shortcut form', 'submit', async (event) => {
        event.preventDefault()

        const modal = $('#modal-add-shortcut')
        const form = $('form', modal)
        const button = $('button[type="submit"]', form)

        if (button.disabled) {
            return
        }

        const data = new FormData(form)

        /** @type {AddShortcutData} */
        const body = {
            id: data.get('id'),
            program: data.get('program'),
            name: data.get('name'),
            description: data.get('description'),
            startDirectory: data.get('startDirectory'),
            executable: data.get('executable'),
            launchOptions: data.get('launchOptions'),
            relativePath: data.get('relativePath'),
            iconUrl: data.get('icon'),
            logoUrl: data.get('logo'),
            coverUrl: data.get('cover'),
            bannerUrl: data.get('banner'),
            heroUrl: data.get('hero'),
            tags: data.get('tags').split(',')
        }

        await window.runAndCaptureConsole(button, false, async () => {
            try {
                /** @type {AddShortcutResult} */
                await requestJson('POST', '/api/shortcut/add', JSON.stringify(body))
                /** @type {SaveLibraryResult} */
                await requestJson('POST', '/api/library/save')
                await loadShortcuts()
            } catch (error) {
                window.showError(error)
            }
        })

        window.hideModal(modal)

    })

    on('#modal-update-shortcut form', 'submit', async (event) => {
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
            relativePath: data.get('relativePath'),
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

    on('#modal-delete-shortcut form', 'submit', async (event) => {
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

    on('#modal-sync-shortcuts form', 'submit', async (event) => {
        event.preventDefault()

        const modal = $('#modal-sync-shortcuts')
        const form = $('form', modal)
        const button = $('button[type="submit"]', form)

        if (button.disabled) {
            return
        }

        await window.runAndCaptureConsole(button, false, async () => {
            try {
                /** @type {LoadLibraryResult} */
                await requestJson('POST', '/api/library/load')

                /** @type {SyncLibraryResult} */
                await requestJson('POST', '/api/library/sync')

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

    window.loadShortcuts = loadShortcuts

})