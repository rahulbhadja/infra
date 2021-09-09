/* tslint:disable */
/* eslint-disable */
/**
 * Infra API
 * Infra REST API
 *
 * The version of the OpenAPI document: 0.1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import {
    Source,
    SourceFromJSON,
    SourceToJSON,
} from '../models';

/**
 * 
 */
export class SourcesApi extends runtime.BaseAPI {

    /**
     * List sources
     */
    async listSourcesRaw(initOverrides?: RequestInit): Promise<runtime.ApiResponse<Array<Source>>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/sources`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(SourceFromJSON));
    }

    /**
     * List sources
     */
    async listSources(initOverrides?: RequestInit): Promise<Array<Source>> {
        const response = await this.listSourcesRaw(initOverrides);
        return await response.value();
    }

}