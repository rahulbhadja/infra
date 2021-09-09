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
    Role,
    RoleFromJSON,
    RoleToJSON,
} from '../models';

export interface ListRolesRequest {
    destinationId: string;
}

/**
 * 
 */
export class RolesApi extends runtime.BaseAPI {

    /**
     * List roles
     */
    async listRolesRaw(requestParameters: ListRolesRequest, initOverrides?: RequestInit): Promise<runtime.ApiResponse<Array<Role>>> {
        if (requestParameters.destinationId === null || requestParameters.destinationId === undefined) {
            throw new runtime.RequiredError('destinationId','Required parameter requestParameters.destinationId was null or undefined when calling listRoles.');
        }

        const queryParameters: any = {};

        if (requestParameters.destinationId !== undefined) {
            queryParameters['destinationId'] = requestParameters.destinationId;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = await token("bearerAuth", []);

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/roles`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(RoleFromJSON));
    }

    /**
     * List roles
     */
    async listRoles(requestParameters: ListRolesRequest, initOverrides?: RequestInit): Promise<Array<Role>> {
        const response = await this.listRolesRaw(requestParameters, initOverrides);
        return await response.value();
    }

}