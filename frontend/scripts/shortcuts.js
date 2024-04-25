// Shortcuts
window.addEventListener('load', async () => {

    /** @type {Shortcut[]} */
    let shortcuts = []

    /**
     * Retrieve shortcut by appId
     * @param {Number} appId
     * @returns {Shortcut}
     */
    function getShortcut(appId) {
        return shortcuts.find((item) => item.appId === Number(appId))
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

            const library = await requestJson('POST', '/api/library/load')

            /** @type {Shortcut[]} */
            shortcuts = await requestJson('GET', '/api/shortcuts')

            const items = shortcuts.map((shortcut) => {
                const coverUrl = (shortcut.cover)
                    ? String(shortcut.cover).replace(library.userArtworksPath, "/grid/image")
                    : './img/default/cover.png'

                return `<article class="item shortcut" title="${shortcut.appName}">
                    <div class="area">
                        <div class="image">
                            <img loading="lazy" src="${coverUrl}" alt="${shortcut.appName}" width="600" height="900"/>
                        </div>
                        <div class="info">
                            <div class="title">
                                <small>${shortcut.appId}</small><br/>
                                <h4>${shortcut.appName}</h4>
                            </div>
                            <div class="actions">
                                <button type="button" data-launch-shortcut="${shortcut.appId}" title="Launch">
                                    <img src="./img/icons/launch.svg" alt="Launch" width="24" height="24" />
                                </button>
                                <button type="button" data-update-shortcut="${shortcut.appId}" title="Update">
                                    <img src="./img/icons/update.svg" alt="Update" width="24" height="24" />
                                </button>
                                <button type="button" data-delete-shortcut="${shortcut.appId}" title="Delete">
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
            }

            const destination = $('#shortcuts #list')
            destination.innerHTML = items.join('')
        } finally {
            button.disabled = false
        }

    }

    on('#shortcuts [data-launch-shortcut]', 'click', async (event) => {
        event.preventDefault()

        const button = event.target.closest('[data-launch-shortcut]')

        if (button.disabled) {
            return
        }

        const modal = $('#modal-launch-shortcut')
        const content = $('.content', modal)
        const shortcut = getShortcut(button.dataset.launchShortcut)
        
        modal.dataset.shortcut = shortcut.appId
        content.innerHTML = `<p>Launching <b>${shortcut.appName}</b>...</p>`
        window.showModal(modal)

        const body = JSON.stringify({
            appId: shortcut.appId
        })

        try {
            button.disabled = true
            await window.runAndCaptureConsole(false, async () => {
                await requestJson('POST', '/api/shortcut/launch', body)
            })
        } catch (error) {
            window.showError(error)
        } finally {
            button.disabled = false
        }

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

        modal.dataset.shortcut = shortcut.appId
        content.innerHTML = '<p>Loading data, please wait...</p>'
        window.showModal(modal)

        try {

            /** @type {ScrapeResult} */
            const scrape = await requestJson('GET', '/api/scrape?term=' + encodeURIComponent(shortcut.appName))

            const html = []
            html.push(
                `<div class="group">
                    <label>Name:</label>
                    <input type="text" name="appName" value="${shortcut.appName}" />
                </div>
                <div class="group">
                    <label>Start Directory:</label>
                    <input type="text" name="startDir" value="${shortcut.startDir}" />
                </div>
                <div class="group">
                    <label>Executable:</label>
                    <input type="text" name="exe" value="${shortcut.exe}" />
                </div>
                <div class="group">
                    <label>Launch Options:</label>
                    <textarea name="launchOptions">${shortcut.launchOptions}</textarea>
                </div>
                `)

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

            append('cover', 'Cover Artworks', shortcut.coverUrl, scrape.result.coverUrls, 600, 900)
            append('banner', 'Banner Artworks', shortcut.bannerUrl, scrape.result.bannerUrls, 920, 430)
            append('hero', 'Hero Artworks', shortcut.heroUrl, scrape.result.heroUrls, 600, 900)
            append('icon', 'Icon Artworks', shortcut.iconUrl, scrape.result.iconUrls, 192, 192)
            append('logo', 'Logo Artworks', shortcut.logoUrl, scrape.result.logoUrls, 600, 900)

            content.innerHTML = html.join('')

        } catch (error) {
            window.showError(error)
            window.hideModal(modal)
        }

    })

    on('#shortcuts [data-delete-shortcut]', 'click', (event) => {

        event.preventDefault()
        const element = event.target.closest('[data-delete-shortcut]')
        const shortcut = getShortcut(element.dataset.deleteShortcut)

        const modal = $('#modal-delete-shortcut')
        const content = $('.content', modal)

        modal.dataset.shortcut = shortcut.appId
        content.innerHTML = `<p>Are you sure you want to delete the shortcut to <b>${shortcut.appName}</b>?</p>`
        window.showModal(modal)

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
        const body = JSON.stringify({
            action: 'update',
            appId: shortcut.appId,
            appName: data.get('appName'),
            startDir: data.get('startDir'),
            exe: data.get('exe'),
            launchOptions: data.get('launchOptions'),
            iconUrl: data.get('icon'),
            logoUrl: data.get('logo'),
            coverUrl: data.get('cover'),
            bannerUrl: data.get('banner'),
            heroUrl: data.get('hero')
        })

        try {
            button.disabled = true
            await window.runAndCaptureConsole(false, async () => {
                await requestJson('POST', '/api/shortcut/modify', body)
                await requestJson('POST', '/api/library/save')
            })
            await loadShortcuts()
        } catch (error) {
            window.showError(error)
        } finally {
            button.disabled = false
        }

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
        const body = JSON.stringify({
            action: 'delete',
            appId: shortcut.appId
        })

        try {
            button.disabled = true
            await window.runAndCaptureConsole(false, async () => {
                await requestJson('POST', '/api/shortcut/modify', body)
                await requestJson('POST', '/api/library/save')
            })
            await loadShortcuts()
        } catch (error) {
            window.showError(error)
        } finally {
            button.disabled = false
        }

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