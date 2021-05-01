"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Events = exports.Commands = exports.Categories = void 0;
var Categories;
(function (Categories) {
    Categories["Users"] = "Users";
    Categories["Project"] = "Project";
    Categories["Estimation"] = "Estimation";
    Categories["Accounting"] = "Accounting";
})(Categories = exports.Categories || (exports.Categories = {}));
var Commands;
(function (Commands) {
    Commands["EnableAccounting"] = "EnableAccounting";
})(Commands = exports.Commands || (exports.Commands = {}));
var Events;
(function (Events) {
    Events["MembershipCreated"] = "MembershipCreated";
    Events["MembershipCreatedForProject"] = "MembershipCreatedForProject";
    Events["MembershipUpdated"] = "MembershipUpdated";
    Events["MembershipDeleted"] = "MembershipDeleted";
    Events["ProjectCreated"] = "ProjectCreated";
    Events["ProjectUpdated"] = "ProjectUpdated";
    Events["ProjectDeleted"] = "ProjectDeleted";
})(Events = exports.Events || (exports.Events = {}));
