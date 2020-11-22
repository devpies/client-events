"use strict";
// System commands and events
Object.defineProperty(exports, "__esModule", { value: true });
exports.Events = exports.Commands = void 0;
// Commands are things to be done. (in the future)
// Events are things that have happened. (in the past)
var Commands;
(function (Commands) {
    Commands["AddUser"] = "AddUser";
    Commands["ModifyUser"] = "ModifyUser";
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
})(Commands = exports.Commands || (exports.Commands = {}));
var Events;
(function (Events) {
    Events["UserAdded"] = "UserAdded";
    Events["UserModified"] = "UserModified";
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
})(Events = exports.Events || (exports.Events = {}));
