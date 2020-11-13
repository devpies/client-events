# To use this code, make sure you
#
#     import json
#
# and then, to convert JSON from a string, do
#
#     result = event_types_from_dict(json.loads(json_string))
#     result = metadata_from_dict(json.loads(json_string))
#     result = registered_user_event_from_dict(json.loads(json_string))
#     result = change_user_event_from_dict(json.loads(json_string))
#     result = user_changed_event_from_dict(json.loads(json_string))

from enum import Enum
from typing import Any, TypeVar, Type, cast


T = TypeVar("T")
EnumT = TypeVar("EnumT", bound=Enum)


def from_str(x: Any) -> str:
    assert isinstance(x, str)
    return x


def from_bool(x: Any) -> bool:
    assert isinstance(x, bool)
    return x


def to_class(c: Type[T], x: Any) -> dict:
    assert isinstance(x, c)
    return cast(Any, x).to_dict()


def to_enum(c: Type[EnumT], x: Any) -> EnumT:
    assert isinstance(x, c)
    return x.value


class EventTypes(Enum):
    CHANGE_USER = "ChangeUser"
    REGISTERED_USER = "RegisteredUser"
    USER_CHANGED = "UserChanged"


class RegisteredUserEventData:
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

    @staticmethod
    def from_dict(obj: Any) -> 'RegisteredUserEventData':
        assert isinstance(obj, dict)
        auth0_id = from_str(obj.get("auth0Id"))
        email = from_str(obj.get("email"))
        email_verified = from_bool(obj.get("emailVerified"))
        first_name = from_str(obj.get("firstName"))
        id = from_str(obj.get("id"))
        last_name = from_str(obj.get("lastName"))
        locale = from_str(obj.get("locale"))
        picture = from_str(obj.get("picture"))
        return RegisteredUserEventData(auth0_id, email, email_verified, first_name, id, last_name, locale, picture)

    def to_dict(self) -> dict:
        result: dict = {}
        result["auth0Id"] = from_str(self.auth0_id)
        result["email"] = from_str(self.email)
        result["emailVerified"] = from_bool(self.email_verified)
        result["firstName"] = from_str(self.first_name)
        result["id"] = from_str(self.id)
        result["lastName"] = from_str(self.last_name)
        result["locale"] = from_str(self.locale)
        result["picture"] = from_str(self.picture)
        return result


class Metadata:
    trace_id: str
    user_id: str

    def __init__(self, trace_id: str, user_id: str) -> None:
        self.trace_id = trace_id
        self.user_id = user_id

    @staticmethod
    def from_dict(obj: Any) -> 'Metadata':
        assert isinstance(obj, dict)
        trace_id = from_str(obj.get("traceId"))
        user_id = from_str(obj.get("userId"))
        return Metadata(trace_id, user_id)

    def to_dict(self) -> dict:
        result: dict = {}
        result["traceId"] = from_str(self.trace_id)
        result["userId"] = from_str(self.user_id)
        return result


class RegisteredUserEventType(Enum):
    REGISTERED_USER = "RegisteredUser"


class RegisteredUserEvent:
    data: RegisteredUserEventData
    id: str
    metadata: Metadata
    type: RegisteredUserEventType

    def __init__(self, data: RegisteredUserEventData, id: str, metadata: Metadata, type: RegisteredUserEventType) -> None:
        self.data = data
        self.id = id
        self.metadata = metadata
        self.type = type

    @staticmethod
    def from_dict(obj: Any) -> 'RegisteredUserEvent':
        assert isinstance(obj, dict)
        data = RegisteredUserEventData.from_dict(obj.get("data"))
        id = from_str(obj.get("id"))
        metadata = Metadata.from_dict(obj.get("metadata"))
        type = RegisteredUserEventType(obj.get("type"))
        return RegisteredUserEvent(data, id, metadata, type)

    def to_dict(self) -> dict:
        result: dict = {}
        result["data"] = to_class(RegisteredUserEventData, self.data)
        result["id"] = from_str(self.id)
        result["metadata"] = to_class(Metadata, self.metadata)
        result["type"] = to_enum(RegisteredUserEventType, self.type)
        return result


class ChangeUserEventData:
    first_name: str
    last_name: str
    locale: str
    picture: str

    def __init__(self, first_name: str, last_name: str, locale: str, picture: str) -> None:
        self.first_name = first_name
        self.last_name = last_name
        self.locale = locale
        self.picture = picture

    @staticmethod
    def from_dict(obj: Any) -> 'ChangeUserEventData':
        assert isinstance(obj, dict)
        first_name = from_str(obj.get("firstName"))
        last_name = from_str(obj.get("lastName"))
        locale = from_str(obj.get("locale"))
        picture = from_str(obj.get("picture"))
        return ChangeUserEventData(first_name, last_name, locale, picture)

    def to_dict(self) -> dict:
        result: dict = {}
        result["firstName"] = from_str(self.first_name)
        result["lastName"] = from_str(self.last_name)
        result["locale"] = from_str(self.locale)
        result["picture"] = from_str(self.picture)
        return result


class ChangeUserEventType(Enum):
    CHANGE_USER = "ChangeUser"


class ChangeUserEvent:
    data: ChangeUserEventData
    id: str
    metadata: Metadata
    type: ChangeUserEventType

    def __init__(self, data: ChangeUserEventData, id: str, metadata: Metadata, type: ChangeUserEventType) -> None:
        self.data = data
        self.id = id
        self.metadata = metadata
        self.type = type

    @staticmethod
    def from_dict(obj: Any) -> 'ChangeUserEvent':
        assert isinstance(obj, dict)
        data = ChangeUserEventData.from_dict(obj.get("data"))
        id = from_str(obj.get("id"))
        metadata = Metadata.from_dict(obj.get("metadata"))
        type = ChangeUserEventType(obj.get("type"))
        return ChangeUserEvent(data, id, metadata, type)

    def to_dict(self) -> dict:
        result: dict = {}
        result["data"] = to_class(ChangeUserEventData, self.data)
        result["id"] = from_str(self.id)
        result["metadata"] = to_class(Metadata, self.metadata)
        result["type"] = to_enum(ChangeUserEventType, self.type)
        return result


class UserChangedEventData:
    first_name: str
    last_name: str
    locale: str
    picture: str

    def __init__(self, first_name: str, last_name: str, locale: str, picture: str) -> None:
        self.first_name = first_name
        self.last_name = last_name
        self.locale = locale
        self.picture = picture

    @staticmethod
    def from_dict(obj: Any) -> 'UserChangedEventData':
        assert isinstance(obj, dict)
        first_name = from_str(obj.get("firstName"))
        last_name = from_str(obj.get("lastName"))
        locale = from_str(obj.get("locale"))
        picture = from_str(obj.get("picture"))
        return UserChangedEventData(first_name, last_name, locale, picture)

    def to_dict(self) -> dict:
        result: dict = {}
        result["firstName"] = from_str(self.first_name)
        result["lastName"] = from_str(self.last_name)
        result["locale"] = from_str(self.locale)
        result["picture"] = from_str(self.picture)
        return result


class UserChangedEventType(Enum):
    USER_CHANGED = "UserChanged"


class UserChangedEvent:
    data: UserChangedEventData
    id: str
    metadata: Metadata
    type: UserChangedEventType

    def __init__(self, data: UserChangedEventData, id: str, metadata: Metadata, type: UserChangedEventType) -> None:
        self.data = data
        self.id = id
        self.metadata = metadata
        self.type = type

    @staticmethod
    def from_dict(obj: Any) -> 'UserChangedEvent':
        assert isinstance(obj, dict)
        data = UserChangedEventData.from_dict(obj.get("data"))
        id = from_str(obj.get("id"))
        metadata = Metadata.from_dict(obj.get("metadata"))
        type = UserChangedEventType(obj.get("type"))
        return UserChangedEvent(data, id, metadata, type)

    def to_dict(self) -> dict:
        result: dict = {}
        result["data"] = to_class(UserChangedEventData, self.data)
        result["id"] = from_str(self.id)
        result["metadata"] = to_class(Metadata, self.metadata)
        result["type"] = to_enum(UserChangedEventType, self.type)
        return result


def event_types_from_dict(s: Any) -> EventTypes:
    return EventTypes(s)


def event_types_to_dict(x: EventTypes) -> Any:
    return to_enum(EventTypes, x)


def metadata_from_dict(s: Any) -> Metadata:
    return Metadata.from_dict(s)


def metadata_to_dict(x: Metadata) -> Any:
    return to_class(Metadata, x)


def registered_user_event_from_dict(s: Any) -> RegisteredUserEvent:
    return RegisteredUserEvent.from_dict(s)


def registered_user_event_to_dict(x: RegisteredUserEvent) -> Any:
    return to_class(RegisteredUserEvent, x)


def change_user_event_from_dict(s: Any) -> ChangeUserEvent:
    return ChangeUserEvent.from_dict(s)


def change_user_event_to_dict(x: ChangeUserEvent) -> Any:
    return to_class(ChangeUserEvent, x)


def user_changed_event_from_dict(s: Any) -> UserChangedEvent:
    return UserChangedEvent.from_dict(s)


def user_changed_event_to_dict(x: UserChangedEvent) -> Any:
    return to_class(UserChangedEvent, x)
