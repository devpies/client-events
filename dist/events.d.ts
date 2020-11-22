export declare enum Commands {
    AddUser = "AddUser",
    ModifyUser = "ModifyUser"
}
export interface AddUserCommand {
    id: string;
    type: Commands.AddUser;
    metadata: Metadata;
    data: {
        id: string;
        auth0Id: string;
        email: string;
        emailVerified: boolean;
        firstName: string;
        lastName: string;
        picture: string;
        locale: string;
    };
}
export interface ModifyUserCommand {
    id: string;
    type: Commands.ModifyUser;
    metadata: Metadata;
    data: {
        firstName: string;
        lastName: string;
        picture: string;
        locale: string;
    };
}
export declare enum Events {
    UserAdded = "UserAdded",
    UserModified = "UserModified"
}
interface Metadata {
    traceId: string;
    userId: string;
}
export interface UserAddedEvent {
    id: string;
    type: Events.UserAdded;
    metadata: Metadata;
    data: {
        id: string;
        auth0Id: string;
        email: string;
        emailVerified: boolean;
        firstName: string;
        lastName: string;
        picture: string;
        locale: string;
    };
}
export interface UserModifiedEvent {
    id: string;
    type: Events.UserModified;
    metadata: Metadata;
    data: {
        firstName: string;
        lastName: string;
        picture: string;
        locale: string;
    };
}
export {};
