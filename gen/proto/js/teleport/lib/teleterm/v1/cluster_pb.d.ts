// package: teleport.lib.teleterm.v1
// file: teleport/lib/teleterm/v1/cluster.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class Cluster extends jspb.Message { 
    getUri(): string;
    setUri(value: string): Cluster;

    getName(): string;
    setName(value: string): Cluster;

    getProxyHost(): string;
    setProxyHost(value: string): Cluster;

    getConnected(): boolean;
    setConnected(value: boolean): Cluster;

    getLeaf(): boolean;
    setLeaf(value: boolean): Cluster;


    hasLoggedInUser(): boolean;
    clearLoggedInUser(): void;
    getLoggedInUser(): LoggedInUser | undefined;
    setLoggedInUser(value?: LoggedInUser): Cluster;


    hasFeatures(): boolean;
    clearFeatures(): void;
    getFeatures(): Features | undefined;
    setFeatures(value?: Features): Cluster;

    getAuthClusterId(): string;
    setAuthClusterId(value: string): Cluster;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Cluster.AsObject;
    static toObject(includeInstance: boolean, msg: Cluster): Cluster.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Cluster, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Cluster;
    static deserializeBinaryFromReader(message: Cluster, reader: jspb.BinaryReader): Cluster;
}

export namespace Cluster {
    export type AsObject = {
        uri: string,
        name: string,
        proxyHost: string,
        connected: boolean,
        leaf: boolean,
        loggedInUser?: LoggedInUser.AsObject,
        features?: Features.AsObject,
        authClusterId: string,
    }
}

export class LoggedInUser extends jspb.Message { 
    getName(): string;
    setName(value: string): LoggedInUser;

    clearRolesList(): void;
    getRolesList(): Array<string>;
    setRolesList(value: Array<string>): LoggedInUser;
    addRoles(value: string, index?: number): string;

    clearSshLoginsList(): void;
    getSshLoginsList(): Array<string>;
    setSshLoginsList(value: Array<string>): LoggedInUser;
    addSshLogins(value: string, index?: number): string;


    hasAcl(): boolean;
    clearAcl(): void;
    getAcl(): ACL | undefined;
    setAcl(value?: ACL): LoggedInUser;

    clearActiveRequestsList(): void;
    getActiveRequestsList(): Array<string>;
    setActiveRequestsList(value: Array<string>): LoggedInUser;
    addActiveRequests(value: string, index?: number): string;

    clearSuggestedReviewersList(): void;
    getSuggestedReviewersList(): Array<string>;
    setSuggestedReviewersList(value: Array<string>): LoggedInUser;
    addSuggestedReviewers(value: string, index?: number): string;

    clearRequestableRolesList(): void;
    getRequestableRolesList(): Array<string>;
    setRequestableRolesList(value: Array<string>): LoggedInUser;
    addRequestableRoles(value: string, index?: number): string;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): LoggedInUser.AsObject;
    static toObject(includeInstance: boolean, msg: LoggedInUser): LoggedInUser.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: LoggedInUser, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): LoggedInUser;
    static deserializeBinaryFromReader(message: LoggedInUser, reader: jspb.BinaryReader): LoggedInUser;
}

export namespace LoggedInUser {
    export type AsObject = {
        name: string,
        rolesList: Array<string>,
        sshLoginsList: Array<string>,
        acl?: ACL.AsObject,
        activeRequestsList: Array<string>,
        suggestedReviewersList: Array<string>,
        requestableRolesList: Array<string>,
    }
}

export class ACL extends jspb.Message { 

    hasAuthConnectors(): boolean;
    clearAuthConnectors(): void;
    getAuthConnectors(): ResourceAccess | undefined;
    setAuthConnectors(value?: ResourceAccess): ACL;


    hasRoles(): boolean;
    clearRoles(): void;
    getRoles(): ResourceAccess | undefined;
    setRoles(value?: ResourceAccess): ACL;


    hasUsers(): boolean;
    clearUsers(): void;
    getUsers(): ResourceAccess | undefined;
    setUsers(value?: ResourceAccess): ACL;


    hasTrustedClusters(): boolean;
    clearTrustedClusters(): void;
    getTrustedClusters(): ResourceAccess | undefined;
    setTrustedClusters(value?: ResourceAccess): ACL;


    hasEvents(): boolean;
    clearEvents(): void;
    getEvents(): ResourceAccess | undefined;
    setEvents(value?: ResourceAccess): ACL;


    hasTokens(): boolean;
    clearTokens(): void;
    getTokens(): ResourceAccess | undefined;
    setTokens(value?: ResourceAccess): ACL;


    hasServers(): boolean;
    clearServers(): void;
    getServers(): ResourceAccess | undefined;
    setServers(value?: ResourceAccess): ACL;


    hasApps(): boolean;
    clearApps(): void;
    getApps(): ResourceAccess | undefined;
    setApps(value?: ResourceAccess): ACL;


    hasDbs(): boolean;
    clearDbs(): void;
    getDbs(): ResourceAccess | undefined;
    setDbs(value?: ResourceAccess): ACL;


    hasKubeservers(): boolean;
    clearKubeservers(): void;
    getKubeservers(): ResourceAccess | undefined;
    setKubeservers(value?: ResourceAccess): ACL;


    hasAccessRequests(): boolean;
    clearAccessRequests(): void;
    getAccessRequests(): ResourceAccess | undefined;
    setAccessRequests(value?: ResourceAccess): ACL;


    hasRecordedSessions(): boolean;
    clearRecordedSessions(): void;
    getRecordedSessions(): ResourceAccess | undefined;
    setRecordedSessions(value?: ResourceAccess): ACL;


    hasActiveSessions(): boolean;
    clearActiveSessions(): void;
    getActiveSessions(): ResourceAccess | undefined;
    setActiveSessions(value?: ResourceAccess): ACL;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ACL.AsObject;
    static toObject(includeInstance: boolean, msg: ACL): ACL.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ACL, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ACL;
    static deserializeBinaryFromReader(message: ACL, reader: jspb.BinaryReader): ACL;
}

export namespace ACL {
    export type AsObject = {
        authConnectors?: ResourceAccess.AsObject,
        roles?: ResourceAccess.AsObject,
        users?: ResourceAccess.AsObject,
        trustedClusters?: ResourceAccess.AsObject,
        events?: ResourceAccess.AsObject,
        tokens?: ResourceAccess.AsObject,
        servers?: ResourceAccess.AsObject,
        apps?: ResourceAccess.AsObject,
        dbs?: ResourceAccess.AsObject,
        kubeservers?: ResourceAccess.AsObject,
        accessRequests?: ResourceAccess.AsObject,
        recordedSessions?: ResourceAccess.AsObject,
        activeSessions?: ResourceAccess.AsObject,
    }
}

export class ResourceAccess extends jspb.Message { 
    getList(): boolean;
    setList(value: boolean): ResourceAccess;

    getRead(): boolean;
    setRead(value: boolean): ResourceAccess;

    getEdit(): boolean;
    setEdit(value: boolean): ResourceAccess;

    getCreate(): boolean;
    setCreate(value: boolean): ResourceAccess;

    getDelete(): boolean;
    setDelete(value: boolean): ResourceAccess;

    getUse(): boolean;
    setUse(value: boolean): ResourceAccess;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ResourceAccess.AsObject;
    static toObject(includeInstance: boolean, msg: ResourceAccess): ResourceAccess.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ResourceAccess, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ResourceAccess;
    static deserializeBinaryFromReader(message: ResourceAccess, reader: jspb.BinaryReader): ResourceAccess;
}

export namespace ResourceAccess {
    export type AsObject = {
        list: boolean,
        read: boolean,
        edit: boolean,
        create: boolean,
        pb_delete: boolean,
        use: boolean,
    }
}

export class Features extends jspb.Message { 
    getAdvancedAccessWorkflows(): boolean;
    setAdvancedAccessWorkflows(value: boolean): Features;

    getIsUsageBasedBilling(): boolean;
    setIsUsageBasedBilling(value: boolean): Features;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Features.AsObject;
    static toObject(includeInstance: boolean, msg: Features): Features.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Features, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Features;
    static deserializeBinaryFromReader(message: Features, reader: jspb.BinaryReader): Features;
}

export namespace Features {
    export type AsObject = {
        advancedAccessWorkflows: boolean,
        isUsageBasedBilling: boolean,
    }
}
