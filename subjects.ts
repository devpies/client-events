// commands and events
// commands are things to be done in the future
// events are thing that have happened

export enum Subjects {
    RegisteredUser = 'RegisteredUser',
    ChangeUser = 'ChangeUser',
    UserChanged = 'UserChanged',
    AddTask = 'AddTask',
    TaskAdded = 'TaskAdded',
    ChangeTask = 'ChangeTask',
    TaskChanged = 'TaskChanged',
    MoveTask = 'MoveTask',
    TaskMoved = 'TaskMoved',
    RemoveTask = 'RemoveTask',
    TaskRemoved = 'TaskRemoved',
    AddColumn = 'AddColumn',
    ColumnAdded = 'ColumnAdded',
    ChangeColumn = 'ChangeColumn',
    ColumnChanged = 'ColumnChanged',
    RemoveColumn = 'RemoveColumn',
    ColumnRemoved = 'ColumnRemoved',
    AddProject = 'AddProject',
    ProjectAdded = 'ProjectAdded',
    ChangeProject = 'ChangeProject',
    ProjectChanged = 'ProjectChanged',
    RemoveProject = 'RemoveProject',
    ProjectRemoved = 'ProjectRemoved'
}