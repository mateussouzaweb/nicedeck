// Shortcuts
window.addEventListener('load', async () => {

    async function loadShortcuts(){

        const button = $('#shortcuts #load')
        button.disabled = true

        try {
            /** @type {Shortcut[]} */
            const shortcuts = await requestJson('GET', '/api/shortcuts')

            const items = shortcuts.map((shortcut) => {
                return `<article class="item" title="${shortcut.appName}">
                    <img loading="lazy" src="${shortcut.coverUrl}" alt="${shortcut.appName}" width="600" height="900"/><br/>
                    <small>${shortcut.appId}</small><br/>
                    <h4>${shortcut.appName}</h4>
                    ${shortcut.platform}
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