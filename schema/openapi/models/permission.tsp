model PagePermissions {
  settings: SettingPagePermissions;
}

model SettingPagePermissions {
  users: boolean;
  roles: boolean;
}
