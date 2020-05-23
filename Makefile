# Copyright 2020 The worker-monitor-client author.

TARGET = worker-monitor.exe

# Definition of commands
GOCMD   = go
GOBUILD = $(GOCMD) build
CD      = powershell.exe cd
CPDIR   = powershell.exe cp -r
MKDIR   = powershell.exe New-Item -ItemType Directory
RM      = powershell.exe Remove-Item -Recurse

# Definition of variables
BINDIR       = bin
SRCDIR       = src
RESULTDIR    = result
CONFIGDIR    = config
CONFIGFILE   = config.dat
INSTALLPATH  = '"C:\Program Files\worker-monitor"'

# Main
$(TARGET) : $(TARGET)
		$(MKDIR) $(BINDIR)
		$(CD) $(SRCDIR); $(GOBUILD) -o ../$(BINDIR)/$(TARGET)

# Task
.PHONY : install
install  :
		$(MKDIR) $(INSTALLPATH)
		$(CPDIR) $(RESULTDIR) $(INSTALLPATH)
		$(CPDIR) $(BINDIR)    $(INSTALLPATH)
		$(CPDIR) $(CONFIGDIR) $(INSTALLPATH)

.PHONY : clean
clean  :
		$(RM) $(BINDIR)
		$(RM) $(INSTALLPATH)