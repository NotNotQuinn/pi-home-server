# Websub server
This is a websub hub, to be used by any and all future projects.

## Needed implementations
Generic implementations:
  - [x] Websub subscriber, publisher and hub library.
  - [ ] Generic service library (p3b & p4a)
    - [ ] Data format creator package (p5b) Encode/Decode with struct tags.
    - Uses one instance of websub subcriber/publisher library.
    - Can be GenericActor or GenericEmitter
  - [ ] Logic center application (p4b & p5a)
Specific implementations:
  - [ ] Chat bot needs to integrate GS and GE for:
    - GS for sending messages
    - GE for chat commands
  - [ ] "Streamdeck" numpad needs to integrate GE for key events
  - [ ] implement audio server that is a GS for audio play requests
  - [ ] implement wallpaper server that is a GS for wallpaper change requests
