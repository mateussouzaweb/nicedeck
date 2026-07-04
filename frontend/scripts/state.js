// State
window.addEventListener('load', async () => {

    /**
     * Load and show available platforms in the software
     */
    async function loadPlatforms() {

        /** @type {ListPlatformsResult} */
        const request = await requestJson('GET', '/api/platforms')
        const platforms = []

        request.console.map((platform) => {
            platforms.push({
                key: platform.console.toLowerCase().replaceAll(' ', '-'),
                name: platform.console,
                value: platform.name
            })
        })

        const options = platforms.map((platform) => {
            return `<label class="checkbox" title="${platform.name}">
                <input type="checkbox" name="platforms[]" value="${platform.value}" checked="checked" />
                <div class="area">
                    <div class="icon">
                        <img loading="lazy" src="/img/platforms/${platform.key}.png" alt="${platform.name}" width="96" height="96" />
                    </div>
                    <div class="info">
                        <b>${platform.value}</b>
                    </div>
                </div>
            </label>`
        })

        // Empty elements for flexbox
        options.push('<div class="fill"></div>')
        options.push('<div class="fill"></div>')
        options.push('<div class="fill"></div>')
        options.push('<div class="fill"></div>')
        options.push('<div class="fill"></div>')
        options.push('<div class="fill"></div>')
        options.push('<div class="fill"></div>')
        options.push('<div class="fill"></div>')
        options.push('<div class="fill"></div>')
        options.push('<div class="fill"></div>')

        const destination = $('#state .platforms')
        destination.innerHTML = options.join('')

    }

    /**
     * Load and show state preview
     */
    async function loadState() {

        const form = $('#state form')
        const data = new FormData(form)
        const action = data.getAll('action')
        const platforms = data.getAll('platforms[]')
        const preferences = data.getAll('preferences[]')

        /** @type {ListStateParams} */
        const params = new URLSearchParams({
            action: action,
            platforms: platforms.join(','),
            preferences: preferences.join(',')
        })

        /** @type {ListStateResult} */
        const request = await requestJson('GET', '/api/state?' + params.toString())
        const state = request.data || []

        const html = []
        html.push('<table class="table">')
        html.push(
            `<tr>
                <th title="Recommended" width="42">#</th>
                <th>Platform</th>
                <th>Source</th>
                <th>Destination</th>
            </tr>`)

        state.map((stateItem) => {
            const platform = stateItem.platform
            const type = stateItem.type
            const recommended = stateItem.recommended
            const source = stateItem.source
            const destination = stateItem.destination

            html.push(
                `<tr>
                <td class="first">
                    ${recommended ? '<input type="checkbox" readonly checked="checked" />' : '<small>-</small>'}
                </td>
                <td>
                    <b>${platform}</b><br/>
                    <small>(${type})</small>
                </td>
                <td>
                    <span>${source.path}</span><br/>
                    <small>Size: ${source.size} bytes ${source.exist ? '' : '(no exist)'}</small><br/>
                    <small>Modified: ${source.modifiedTime}</small>
                </td>
                <td>
                    <span>${destination.path}</span><br/>
                    <small>Size: ${destination.size} bytes ${destination.exist ? '' : '(no exist)'}</small><br/>
                    <small>Modified: ${destination.modifiedTime}</small>
                </td>
            </tr>`)
        })

        if (state.length == 0) {
            html.push(
                `<tr>
                <td class="first empty" colspan="4">
                    No valid state data detected to perform action.
                </td>
            </tr>`)
        }

        html.push('</table>')

        const destination = $('#state .list')
        destination.innerHTML = html.join('')

    }

    on('#state .select-all', 'click', (event) => {
        event.preventDefault()
        const parent = event.target.closest('.group')
        const inputs = $$('input[type="checkbox"]', parent)
        inputs.map((input) => input.checked = "checked")
        loadState()
    })

    on('#state .clear-all', 'click', (event) => {
        event.preventDefault()
        const parent = event.target.closest('.group')
        const inputs = $$('input[type="checkbox"]', parent)
        inputs.map((input) => input.checked = "")
        loadState()
    })

    on('#state .radio input, #state .checkbox input', 'change', () => {
        loadState()
    })

    on('#state form', 'submit', async (event) => {
        event.preventDefault()

        const form = $('#state form')
        const button = $('button[type="submit"]', form)

        if (button.disabled) {
            return
        }

        const data = new FormData(form)
        const action = data.getAll('action')
        const platforms = data.getAll('platforms[]')
        const preferences = data.getAll('preferences[]')

        /** @type {SyncStateData} */
        const body = {
            platforms: platforms,
            preferences: preferences
        }

        await window.runAndCaptureConsole(button, true, async () => {
            try {
                /** @type {BackupStateResult|RestoreStateResult} */
                await requestJson('POST', `/api/state/${action}`, JSON.stringify(body))
                await loadState()
            } catch (error) {
                window.showError(error)
            }
        })
    })

    try {
        await loadPlatforms()
        await loadState()
    } catch (error) {
        window.showError(error)
    }

})