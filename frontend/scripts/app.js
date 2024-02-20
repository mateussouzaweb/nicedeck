/**
 * Retrieve element with given selector
 * @param {String} selector 
 * @param {Element} parent 
 * @returns 
 */
function $(selector, parent) {
    if (parent) {
        return parent.querySelector(selector)
    }
    return document.querySelector(selector)
}

/**
 * Retrieve elements with given selector
 * @param {String} selector 
 * @param {Element} parent 
 * @returns {Element[]}
 */
function $$(selector, parent) {

    const items = []
    const found = (parent)
        ? parent.querySelectorAll(selector)
        : document.querySelectorAll(selector)

    if (found instanceof NodeList) {
        Array.prototype.forEach.call(found, (item) => {
            items.push(item)
        })
    }

    return items
}

/**
 * Attach event listener on element with selector
 * All events are delegated from document 
 * @param {String} selector 
 * @param {String} event 
 * @param {Function} callback 
 */
function on(selector, event, callback) {
    document.addEventListener(event, async (event) => {
        const target = (event.target).closest(selector)
        if (target) {
            callback.apply(target, [event, target])
        }
    }, false)
}

/**
 * Make API request
 * @param {String} method 
 * @param {String} endpoint 
 * @param {FormData} data 
 * @returns 
 */
async function request(method, endpoint, data) {
    const response = await fetch(endpoint, {
        method: method,
        body: data
    })

    const result = await response.text()
    return result
}

/**
 * Make API request and return JSON output
 * @param {String} method 
 * @param {String} endpoint 
 * @param {FormData} data 
 * @returns 
 */
async function requestJson(method, endpoint, data) {
    const result = await request(method, endpoint, data)
    const json = JSON.parse(result || '{}')

    if (json.error) {
        throw new Error(json.error);
    }

    return json
}