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

%files
/usr/bin/line_notify

%changelog
# let skip this for now
