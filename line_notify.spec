Name:       line_notify
Version:    1%{?dist}
Release:    1
Summary:    Line Notify Package
License:    FIXME

BuildRoot: %{_tmppath}/%{name}-%{version}-buildroot

%description
Line Notify Package

%prep
rm -rf $RPM_BUILD_ROOT

%build

%install
mkdir -p %{buildroot}/usr/bin/
install -m 755 /usr/bin/line_notify %{buildroot}/usr/bin/line_notify
install -m 755 /usr/etc/bible_jp.json %{buildroot}/usr/etc/bible_jp.json

%files
/usr/bin/line_notify
/usr/etc/bible_jp.json

%changelog
# let skip this for now
