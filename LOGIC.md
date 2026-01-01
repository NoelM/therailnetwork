# The Rail Network

The Rail Network's reference documentation of its inner logic.

## Signals

Signals are safety devices which ensure a physical separation between trains.
They convey information to the train driver so that driving behavior can be adjusted accordingly

Signals are capable of displaying multiple statuses–at least 2. Usually, they display one status
at a time, but more rarely they may display a combination of statuses. Their 2 basic states are:
* Open: allowing trains to cross the signal, restrictions may apply.
* Closed: prohibiting trains to cross the signal.

The signals convey their status throught multiple ways:
* Lights, e.g. green, orange, red.
* Radio, for in-cabin annouces.
* Physically, e.g. displaying sign or not.
* A combination of above.

The simulator handles two trigger regimes:
* Manual: applies to stations, bifurcations, deposits or lines with very low traffic.
* Automatic: applies to main lines.

The trigger regime corresponds to the method used to open the signal. For manual signals,
they can be connected to other signals and automatically adapt thanks to other signal status. However,
they cannot be opened without a manual action. 


Signals own either a track circuit or awxle counters. Signals may have two directions, thus, they are
related one to another in order to manager against traffic.

### Signal Propeties

- Handle against traffic, has a default status when not "triggered"
- Should we link two signals if against traffic?

- Position: trackId, pk
- Regime: manual, automaic
- Type: BAL, BAPR, TVM...
- Directions: Ascending, Descending, Both
- Parent ID: when connected to another, this may be changed regarding the node statuses
- Detector: axle counter, track circuit
- Crossable: true, false
- Node: which node relates to it
- Node Open Position: when the signal is opened regarding the a defined node position

### Manual systems

- Closed by default in stations
- Closed by default in transversal direction when arriving at 
