from enum import Enum


class Commands(Enum):
    MODIFY_USER = "ModifyUser"
    REGISTER_USER = "RegisterUser"


class RegisterUserCommandData:
    auth0_id: str
    email: str
    email_verified: bool
    first_name: str
    id: str
    last_name: str
    locale: str
    picture: str

    def __init__(self, auth0_id: str, email: str, email_verified: bool, first_name: str, id: str, last_name: str, locale: str, picture: str) -> None:
        self.auth0_id = auth0_id
        self.email = email
        self.email_verified = email_verified
        self.first_name = first_name
        self.id = id
        self.last_name = last_name
        self.locale = locale
        self.picture = picture


class Metadata:
    trace_id: str
    user_id: str

    def __init__(self, trace_id: str, user_id: str) -> None:
        self.trace_id = trace_id
        self.user_id = user_id


class RegisterUserCommandType(Enum):
    REGISTER_USER = "RegisterUser"


class RegisterUserCommand:
    data: RegisterUserCommandData
    id: str
    metadata: Metadata
    type: RegisterUserCommandType

    def __init__(self, data: RegisterUserCommandData, id: str, metadata: Metadata, type: RegisterUserCommandType) -> None:
        self.data = data
        self.id = id
        self.metadata = metadata
        self.type = type


class ModifyUserCommandData:
    first_name: str
    last_name: str
    locale: str
    picture: str

    def __init__(self, first_name: str, last_name: str, locale: str, picture: str) -> None:
        self.first_name = first_name
        self.last_name = last_name
        self.locale = locale
        self.picture = picture


class ModifyUserCommandType(Enum):
    MODIFY_USER = "ModifyUser"


class ModifyUserCommand:
    data: ModifyUserCommandData
    id: str
    metadata: Metadata
    type: ModifyUserCommandType

    def __init__(self, data: ModifyUserCommandData, id: str, metadata: Metadata, type: ModifyUserCommandType) -> None:
        self.data = data
        self.id = id
        self.metadata = metadata
        self.type = type


class Events(Enum):
    USER_MODIFIED = "UserModified"
    USER_REGISTERED = "UserRegistered"


class UserRegisteredEventData:
    auth0_id: str
    email: str
    email_verified: bool
    first_name: str
    id: str
    last_name: str
    locale: str
    picture: str

    def __init__(self, auth0_id: str, email: str, email_verified: bool, first_name: str, id: str, last_name: str, locale: str, picture: str) -> None:
        self.auth0_id = auth0_id
        self.email = email
        self.email_verified = email_verified
        self.first_name = first_name
        self.id = id
        self.last_name = last_name
        self.locale = locale
        self.picture = picture


class UserRegisteredEventType(Enum):
    USER_REGISTERED = "UserRegistered"


class UserRegisteredEvent:
    data: UserRegisteredEventData
    id: str
    metadata: Metadata
    type: UserRegisteredEventType

    def __init__(self, data: UserRegisteredEventData, id: str, metadata: Metadata, type: UserRegisteredEventType) -> None:
        self.data = data
        self.id = id
        self.metadata = metadata
        self.type = type


class UserModifiedEventData:
    first_name: str
    last_name: str
    locale: str
    picture: str

    def __init__(self, first_name: str, last_name: str, locale: str, picture: str) -> None:
        self.first_name = first_name
        self.last_name = last_name
        self.locale = locale
        self.picture = picture


class UserModifiedEventType(Enum):
    USER_MODIFIED = "UserModified"


class UserModifiedEvent:
    data: UserModifiedEventData
    id: str
    metadata: Metadata
    type: UserModifiedEventType

    def __init__(self, data: UserModifiedEventData, id: str, metadata: Metadata, type: UserModifiedEventType) -> None:
        self.data = data
        self.id = id
        self.metadata = metadata
        self.type = type
