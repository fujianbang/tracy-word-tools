import axios from './axios'
import _pick from 'lodash/pick'
import _assign from 'lodash/assign'
import _isEmpty from 'lodash/isEmpty'

import { API_DEFAULT_CONFIG } from '@/config/index'
import API_CONFIG from '@/service/api'


class MakeApi {
    constructor(options) {
        this.api = {}
        this.apiBuilder(options)
    }

    apiBuilder({
    	sep = '|',
    	config = {},
    	mock = false,
    	debug = false,
		prodBaseURL = '',
    	mockBaseURL = ''
    }) {
    	Object.keys(config).map(namespace => {
    		this._apiSingleBuilder({
                namespace,
                mock,
                mockBaseURL,
                prodBaseURL,
                sep, debug, config: config[namespace]
            })
    	})
    }
    _apiSingleBuilder({
    	namespace,
    	sep = '|',
    	config = {},
    	mock = false,
    	debug = false,
    	prodBaseURL = '',
    	mockBaseURL = ''
    }) {
        config.forEach( api => {
            const {name, desc, params, method, path, mockPath } = api
            let apiname = `${namespace}${sep}${name}`,
                url = mock ? mockPath : path,
                baseURL = mock ? mockBaseURL : prodBaseURL

            Object.defineProperty(this.api, `${namespace}${sep}${name}`, {
                value(outerParams, outerOptions) {
                    let _data = _isEmpty(outerParams) ? params : _pick(_assign({}, params, outerParams), Object.keys(params))
                    return axios(_normoalize(_assign({
                        url,
                        desc,
                        baseURL,
                        method
                    }, outerOptions), _data))
                }
            })
        })
    }
}

function _normoalize(options, data) {
    if (options.method.toUpperCase() === 'POST') {
        options.data = data
    } else if (options.method.toUpperCase() === 'GET') {
        options.params = data
    }
    return options
}

export default new MakeApi({
	config: API_CONFIG,
	...API_DEFAULT_CONFIG
})['api']

