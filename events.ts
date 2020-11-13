// commands and events
// commands are things to be done in the future
// events are thing that have happened

export enum EventTypes {
    RegisteredUser = 'RegisteredUser',
    ChangeUser = 'ChangeUser',
    UserChanged = 'UserChanged',
    // AddTask = 'AddTask',
    // TaskAdded = 'TaskAdded',
    // ChangeTask = 'ChangeTask',
    // TaskChanged = 'TaskChanged',
    // MoveTask = 'MoveTask',
    // TaskMoved = 'TaskMoved',
    // RemoveTask = 'RemoveTask',
    // TaskRemoved = 'TaskRemoved',
    // AddColumn = 'AddColumn',
    // ColumnAdded = 'ColumnAdded',
    // ChangeColumn = 'ChangeColumn',
    // ColumnChanged = 'ColumnChanged',
    // RemoveColumn = 'RemoveColumn',
    // ColumnRemoved = 'ColumnRemoved',
    // AddProject = 'AddProject',
    // ProjectAdded = 'ProjectAdded',
    // ChangeProject = 'ChangeProject',
    // ProjectChanged = 'ProjectChanged',
    // RemoveProject = 'RemoveProject',
    // ProjectRemoved = 'ProjectRemoved'
}

interface Metadata {
    traceId: string;
    userId: string;
}

export interface RegisteredUserEvent {
    id: string;
    type: EventTypes.RegisteredUser;
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

export interface ChangeUserEvent {
    id: string;
    type: EventTypes.ChangeUser;
    metadata: Metadata;
    data: {
        firstName: string;
        lastName: string;
        picture: string;
        locale: string;
    }
}

export interface UserChangedEvent {
    id: string;
    type: EventTypes.UserChanged;
    metadata: Metadata;
    data: {
        firstName: string;
        lastName: string;
        picture: string;
        locale: string;
    }
}