`Definition of Transmission Control Protocol` 

    Connection oriented, reliable and stream based protocol

![TCP life cycle](tlc.png)


- Connect (Three handshakes) 

    1. Client sends a SYN package to the server, transits into SYN_SEND state.
    
    2. Server receives the SYN witha SEQ number, then replies with a ACK, SYN along with a SEQ + 1, then transits into SYN_RECEIVED state

    3. Client receives the package, check the SEQ number and returns with a ACK if it matches, then transits into a ESTABLISHED state.


    There are a few things in terms of the three handshakes. Normally, in order to establish a tcp connection, those three handshakes are required. First of all, the server side actually gives the power of actively establish a connection to the client, since the client side has the entire context of which connection should be established, or say, which file descriptor should be used. Image this case, if there's a traffic congestion, the server receives the SYN + SEQ package where the sequence number is already expired at the client side. If server decides to establish connection based on that sequence number, then error will occur. Also there are cases of using only two handshakes in order to establish the connection, which is called TCP Fast Open, it basically let the application to solve the so called "History Connection Issue" 


- Close (Four wave hands)

    1. Client sends a FIN package that indicates that there's no data it wants to send   so please shutdown the connection elegantly, and transits into FIN_WAIT_1 state
    2. When server side receives the FIN package, it will transit into CLOSE_WAIT state, and send a ACK package to the client. After client receives the ACK, it transits into a FIND_WATI_2 state.
    3. When there's no data needed to be sent at the server side, it will send a FIN package to the client. Client receives the FIN package, replies with a ACK package and transits into TIME_WAIT state, server receives the ACK and transits into CLOSED state.
    4. The client will wait for 2 Maximum segment lifetime, then it will transit into CLOSED state


