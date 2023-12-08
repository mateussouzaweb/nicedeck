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
        button.disabled = true

        try {
            /** @type {Shortcut[]} */
            const library = await requestJson('POST', '/api/library/load')
            shortcuts = await requestJson('GET', '/api/shortcuts')

            const items = shortcuts.map((shortcut) => {
                const coverUrl = String(shortcut.cover).replace(library.userArtworksPath, "/grid/image")
                return `<article class="item shortcut" title="${shortcut.appName}">
                    <div class="area">
                        <div class="image">
                            <img loading="lazy" src="${coverUrl}" alt="${shortcut.appName}" width="600" height="900"/>
                        </div>
                        <div class="info">
                            <small>${shortcut.appId}</small><br/>
                            <h4>${shortcut.appName}</h4>
                            ${shortcut.platform}

                            <div class="actions">
                                <button type="button" data-update-shortcut="${shortcut.appId}">Update</button>
                                <button type="button" data-delete-shortcut="${shortcut.appId}">Delete</button>
                            </div>
                        </div>
                    </div>
                </article>`
            })

            if (!items.length) {
                items.push(`<article class="item">
                    <p>No shortcuts to show here yet...</p>
                </article>`)
            }

            const destination = $('#shortcuts #list')
            destination.innerHTML = items.join('')
        } finally {
            button.disabled = false
        }
    }

    on('#shortcuts [data-update-shortcut]', 'click', async (event) => {

        event.preventDefault()
        const element = event.target.closest('[data-update-shortcut]')
        const shortcut = getShortcut(element.dataset.updateShortcut)

        const modal = $('#modal-update-shortcut')
        const content = $('.content', modal)

        modal.dataset.shortcut = shortcut.appId
        content.innerHTML = '<p>Loading data, please wait...</p>'
        window.showModal(modal)

        /** @type {ScrapeResult} */
        const result = await requestJson('GET', '/api/scrape?term=' + encodeURIComponent(shortcut.appName))
        const html = []

        const append = (type, title, images, width, height) => {
            html.push(`<div class="group">`)
            html.push(`<h4>${title}</h4>`)

            if (!images || !images.length) {
                html.push(`<p>No images found.</p>`)
            } else {
                images.forEach((item, index) => {
                    const checked = index === 0 ? 'checked="checked"' : ''
                    html.push(`
                    <label class="radio">
                        <input type="radio" name="${type}" value="${item}" ${checked} />
                        <div class="image">
                            <img loading="lazy" src="${item}" alt="Image ${index}" 
                            width="${width}" height="${height}"/>
                        </div>
                    </label>
                    `)
                })
            }

            html.push(`</div>`)
        }

        html.push(`<p>Scrape results for <b>${shortcut.appName}</b>:</p>`)
        append('cover', 'Cover Artworks', result.coverUrls, 600, 900)
        append('banner', 'Banner Artworks', result.bannerUrls, 920, 430)
        append('icon', 'Icon Artworks', result.iconUrls, 192, 192)
        append('hero', 'Hero Artworks', result.heroUrls, 600, 900)
        append('logo', 'Logo Artworks', result.logoUrls, 600, 900)
        content.innerHTML = html.join('')

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
        const button = $('button[type="submit"]', modal)
        const form = $('form', modal)

        const shortcut = getShortcut(modal.dataset.shortcut)
        const data = new FormData(form)
        const body = JSON.stringify({
            action: 'update',
            appId: shortcut.appId,
            iconUrl: data.get('icon'),
            logoUrl: data.get('logo'),
            coverUrl: data.get('cover'),
            bannerUrl: data.get('banner'),
            heroUrl: data.get('hero')
        })

        try {
            button.disabled = true
            await window.runAndCaptureConsole(async () => {
                await request('POST', '/api/shortcut/modify', body)
                await request('POST', '/api/library/save')
            })
            await loadShortcuts()
        } finally {
            button.disabled = false
        }

        window.hideModal(modal)

    })

    on('#shortcuts #modal-delete-shortcut form', 'submit', async (event) => {

        event.preventDefault();

        const modal = $('#modal-delete-shortcut')
        const button = $('button[type="submit"]', modal)

        const shortcut = getShortcut(modal.dataset.shortcut)
        const body = JSON.stringify({
            action: 'delete',
            appId: shortcut.appId
        })

        try {
            button.disabled = true
            await window.runAndCaptureConsole(async () => {
                await request('POST', '/api/shortcut/modify', body)
                await request('POST', '/api/library/save')
            })
            await loadShortcuts()
        } finally {
            button.disabled = false
        }

        window.hideModal(modal)

    })

    on('#shortcuts #load', 'click', async (event) => {
        event.preventDefault()
        await loadShortcuts()
    })

    await loadShortcuts()

})