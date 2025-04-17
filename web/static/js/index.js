userIdCookie := &http.Cookie{
    Name:       "session",
    Value:      unevaleur,
    Path:       "",
    Domain:     "",
    Expires:    time.Time{},
    RawExpires: "",
    MaxAge:     99999999999,
    Secure:     false,
    HttpOnly:   false,
    SameSite:   0,
    Raw:        "",
    Unparsed:   []string{},
}