# Copyright 2020 The worker-monitor-client author.

TARGET = worker-monitor.exe

# Definition of commands
GOCMD   = go
GOBUILD = $(GOCMD) build
CD      = powershell.exe cd
CP      = powershell.exe cp
MKDIR   = powershell.exe New-Item -ItemType Directory
RM      = powershell.exe Remove-Item -Recurse

# Definition of variables
BINDIR      = bin
SRCDIR      = src

# Main
$(TARGET) : $(TARGET)
		$(MKDIR) $(BINDIR)
		$(CD) $(SRCDIR); $(GOBUILD) -o ../$(BINDIR)/$(TARGET)

# Task
.PHONY : clean
clean  :
		$(RM) $(BINDIR)