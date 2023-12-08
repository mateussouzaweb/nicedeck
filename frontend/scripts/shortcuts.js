// Shortcuts
window.addEventListener('load', async () => {

    /**
     * Load and show current list of user shortcuts
     */
    async function loadShortcuts(){

        const button = $('#shortcuts #load')
        button.disabled = true

        try {
            /** @type {Shortcut[]} */
            const library = await requestJson('POST', '/api/library/load')
            const shortcuts = await requestJson('GET', '/api/shortcuts')

            const items = shortcuts.map((shortcut) => {
                const coverUrl = String(shortcut.cover).replace(library.userArtworksPath, "/grid/image")
                return `<article class="item" title="${shortcut.appName}">
                    <div class="area">
                        <div class="image">
                            <img loading="lazy" src="${coverUrl}" alt="${shortcut.appName}" width="600" height="900"/>
                        </div>
                        <div class="info">
                            <small>${shortcut.appId}</small><br/>
                            <h4>${shortcut.appName}</h4>
                            ${shortcut.platform}
                        </div>
                    </div>
                </article>`
            })
    
            if( !items.length ){
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

    on('#shortcuts #load', 'click', async (event) => {
        event.preventDefault()
        await loadShortcuts()
    })

    await loadShortcuts()

})