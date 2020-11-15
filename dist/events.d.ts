export declare enum Commands {
    RegisterUser = "RegisterUser",
    ModifyUser = "ModifyUser"
}
export interface RegisterUserCommand {
    id: string;
    type: Commands.RegisterUser;
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
    UserRegistered = "UserRegistered",
    UserModified = "UserModified"
}
interface Metadata {
    traceId: string;
    userId: string;
}
export interface UserRegisteredEvent {
    id: string;
    type: Events.UserRegistered;
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
