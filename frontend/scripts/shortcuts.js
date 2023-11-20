// Shortcuts
window.addEventListener('load', async () => {

    async function loadShortcuts(){

        /** @type {Shortcut[]} */
        const shortcuts = await requestJson('GET', '/api/shortcuts')
        const items = shortcuts.map((shortcut) => {
            return `<article class="item">
                <img loading="lazy" src="${shortcut.coverUrl}" alt="${shortcut.appName}" width="600" height="900"/><br/>
                <small>${shortcut.appId}</small><br/>
                <h4>${shortcut.appName}</h4>
                ${shortcut.platform}
                ${shortcut.platform ? '<br/>' : ''}
            </article>`
        })

        const destination = $('#shortcuts #list')
        destination.innerHTML = items.join('')

    }

    on('#shortcuts #load', 'click', async (event) => {
        event.preventDefault()
        await loadShortcuts()
    })

    await loadShortcuts()

})