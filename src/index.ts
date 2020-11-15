// System commands and events

// Commands are things to be done. (in the future)
// Events are things that have happened. (in the past)

export enum Commands {
    RegisterUser = 'RegisterUser',
    ModifyUser = 'ModifyUser',
    // AddTask = 'AddTask',
    // ModifyTask = 'ModifyTask',
    // DestroyTask = 'DestroyTask',
    // MoveTask = 'MoveTask',
    // AddColumn = 'AddColumn',
    // ModifyColumn = 'ModifyColumn',
    // DestroyColumn = 'DestroyColumn',
    // AddProject = 'AddProject',
    // ModifyProject = 'ModifyProject',
    // DestroyProject = 'DestroyProject',
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
    }
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
    }
}

export enum Events {
    UserRegistered = 'UserRegistered',
    UserModified = 'UserModified',
    // TaskAdded = 'TaskAdded',
    // TaskModified = 'TaskModified',
    // TaskDestroyed = 'TaskDestroyed',
    // TaskMoved = 'TaskMoved',
    // ColumnAdded = 'ColumnAdded',
    // ColumnModified= 'ColumnModified',
    // ColumnDestroyed = 'ColumnDestroyed',
    // ProjectAdded = 'ProjectAdded',
    // ProjectModified= 'ProjectModified',
    // ProjectDestroyed = 'ProjectDestroyed'
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
    }
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
    }
}
