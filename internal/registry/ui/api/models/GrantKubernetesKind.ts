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

/**
 * 
 * @export
 * @enum {string}
 */
export enum GrantKubernetesKind {
    Role = 'role',
    ClusterRole = 'cluster-role'
}

export function GrantKubernetesKindFromJSON(json: any): GrantKubernetesKind {
    return GrantKubernetesKindFromJSONTyped(json, false);
}

export function GrantKubernetesKindFromJSONTyped(json: any, ignoreDiscriminator: boolean): GrantKubernetesKind {
    return json as GrantKubernetesKind;
}

export function GrantKubernetesKindToJSON(value?: GrantKubernetesKind | null): any {
    return value as any;
}
