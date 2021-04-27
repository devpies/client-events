export type EventTypes = Events | Commands;

export interface Message {
  id: string;
  type: EventTypes;
  metadata: Metadata;
  data: any;
}

export interface Metadata {
  traceId: string;
  userId: string;
}

export enum Categories {
  Users = "Users",
  Project = "Project",
  Estimation = "Estimation",
  Accounting = "Accounting",
}

export enum Commands {
  EnableAccounting = "EnableAccounting",
}

export interface EnableAccountingCommand {
  id: string;
  type: Commands.EnableAccounting;
  metadata: Metadata;
  data: {
    auth0Id: string;
    token: string;
  };
}

export enum Events {
  MembershipAdded = "MembershipAdded",
  MembershipDropped = "MembershipDropped",
  ProjectTeamAssigned = "ProjectTeamAssigned",
  ProjectTeamUnassigned = "ProjectTeamUnassigned",
}

export interface MembershipAddedEvent {
  id: string;
  type: Events.MembershipAdded;
  metadata: Metadata;
  data: {
    MemberId: string;
    TeamId: string;
    UserId: string;
    Role: string;
    UpdatedAt: string;
    CreatedAt: string;
  };
}

export interface MembershipDroppedEvent {
  id: string;
  type: Events.MembershipDropped;
  metadata: Metadata;
  data: {
    MemberId: string;
  };
}

export interface ProjectTeamAssignedEvent {
  id: string;
  type: Events.ProjectTeamAssigned;
  metadata: Metadata;
  data: {
    ProjectId: string;
    TeamId: string;
  };
}

export interface ProjectTeamUnassignedEvent {
  id: string;
  type: Events.ProjectTeamUnassigned;
  metadata: Metadata;
  data: {
    ProjectId: string;
  };
}
