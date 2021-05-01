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
  MembershipCreated = "MembershipCreated",
  MembershipCreatedForProject = "MembershipCreatedForProject",
  MembershipUpdated = "MembershipUpdated",
  MembershipDeleted = "MembershipDeleted",
  ProjectCreated = "ProjectCreated",
  ProjectUpdated = "ProjectUpdated",
  ProjectDeleted = "ProjectDeleted",
}

export interface MembershipCreatedEvent {
  id: string;
  type: Events.MembershipCreated;
  metadata: Metadata;
  data: {
    membershipId: string;
    teamId: string;
    userId: string;
    role: string;
    updatedAt: string;
    createdAt: string;
  };
}

export interface MembershipCreatedForProjectEvent {
  id: string;
  type: Events.MembershipCreatedForProject;
  metadata: Metadata;
  data: {
    projectId: string;
    membershipId: string;
    teamId: string;
    userId: string;
    role: string;
    updatedAt: string;
    createdAt: string;
  };
}

export interface MembershipUpdatedEvent {
  id: string;
  type: Events.MembershipUpdated;
  metadata: Metadata;
  data: {
    membershipId: string;
    role: string;
    updatedAt: string;
  };
}

export interface MembershipDeletedEvent {
  id: string;
  type: Events.MembershipDeleted;
  metadata: Metadata;
  data: {
    membershipId: string;
  };
}

export interface ProjectCreatedEvent {
  id: string;
  type: Events.ProjectCreated;
  metadata: Metadata;
  data: {
    projectId: string;
    name: string;
    prefix: string;
    description: string;
    userId: string;
    teamId: string;
    active: boolean;
    public: boolean;
    columnOrder: string[];
    updatedAt: string;
    createdAt: string;
  };
}

export interface ProjectUpdatedEvent {
  id: string;
  type: Events.ProjectUpdated;
  metadata: Metadata;
  data: {
    projectId: string;
    name?: string;
    description?: string;
    teamId?: string;
    active?: boolean;
    public?: boolean;
    columnOrder?: string[];
    updatedAt: string;
  };
}

export interface ProjectDeletedEvent {
  id: string;
  type: Events.ProjectDeleted;
  metadata: Metadata;
  data: {
    projectId: string;
  };
}
