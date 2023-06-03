## Audio virtual output device

create a virtual audio output device

```bash
pactl load-module module-null-sink sink_name=Source
pactl load-module module-virtual-source source_name=VirtualMic master=Source.monitor
```

The first command sets up a “null sink”, i.e. a virtual output device, to collect the audio output from OBS. 
The second command sets up a “virtual source”, i.e. a virtual microphone, and connects it to the monitor channel of the null sink.

## default.pa pulseaudio config

We can automatically load the null sink and virtual source using the default.pa pulseaudio config file

Create the default.pa file 

```bash
touch ~/.config/pulse/default.pa
```

Then add the following code to the default.pa file and save it and then reboot

```bash conf
# include the default.pa pulseaudio config file
.include /etc/pulse/default.pa

# null sink
.ifexists module-null-sink.so
load-module module-null-sink sink_name=Source
.endif

# virtual source
.ifexists module-virtual-source.so
load-module module-virtual-source source_name=VirtualMic master=Source.monitor
.endif
```

Restart PulseAudio
```bash 
pulseaudio -k
```

This will include the /etc/pulse/default.pa
which is the system default pulseaudio and load the null sink and virtual source

<hr>

### Refrences:

- https://youtu.be/Goeucg7A9qE

- https://github.com/NapoleonWils0n/cerberus/blob/master/pulseaudio/virtual-mic.org

- https://github.com/NapoleonWils0n/ubuntu-dotfiles/blob/master/.config/pulse/default.pa

<hr>