# Objects

### About objects

[Objects](https://graphql.github.io/graphql-spec/June2018/#sec-Objects) in GraphQL represent the resources you can access. An object can contain a list of fields, which are specifically typed.

### PageInfo

<p>Information about pagination in a connection.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>hasNextPage</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>When paginating forwards, are there more items?</p></td>
  </tr>
  <tr>
    <td><strong>hasPreviousPage</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>When paginating backwards, are there more items?</p></td>
  </tr>
  <tr>
    <td><strong>startCursor</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>When paginating backwards, the cursor to continue.</p></td>
  </tr>
  <tr>
    <td><strong>endCursor</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>When paginating forwards, the cursor to continue.</p></td>
  </tr>
</table>

---

### SpeciesFilmsConnection

<p>A connection to a list of items.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#speciesfilmsedge">[SpeciesFilmsEdge]</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>totalCount</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>A count of the total number of objects in this connection, ignoring pagination.
This allows a client to fetch the first five objects by passing &ldquo;5&rdquo; as the
argument to &ldquo;first&rdquo;, then fetch the total count so it could display &ldquo;5 of 83&rdquo;,
for example.</p></td>
  </tr>
  <tr>
    <td><strong>films</strong> (<a href="objects.md#film">[Film]</a>)</td> 
    <td><p>A list of all of the objects returned in the connection. This is a convenience
field provided for quickly exploring the API; rather than querying for
&ldquo;{ edges { node } }&rdquo; when no edge data is needed, this field can be be used
instead. Note that when clients like Relay need to fetch the &ldquo;cursor&rdquo; field on
the edge to enable efficient pagination, this shortcut cannot be used, and the
full &ldquo;{ edges { node } }&rdquo; version should be used instead.</p></td>
  </tr>
</table>

---

### PersonFilmsConnection

<p>A connection to a list of items.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#personfilmsedge">[PersonFilmsEdge]</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>totalCount</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>A count of the total number of objects in this connection, ignoring pagination.
This allows a client to fetch the first five objects by passing &ldquo;5&rdquo; as the
argument to &ldquo;first&rdquo;, then fetch the total count so it could display &ldquo;5 of 83&rdquo;,
for example.</p></td>
  </tr>
  <tr>
    <td><strong>films</strong> (<a href="objects.md#film">[Film]</a>)</td> 
    <td><p>A list of all of the objects returned in the connection. This is a convenience
field provided for quickly exploring the API; rather than querying for
&ldquo;{ edges { node } }&rdquo; when no edge data is needed, this field can be be used
instead. Note that when clients like Relay need to fetch the &ldquo;cursor&rdquo; field on
the edge to enable efficient pagination, this shortcut cannot be used, and the
full &ldquo;{ edges { node } }&rdquo; version should be used instead.</p></td>
  </tr>
</table>

---

### SpeciesPeopleEdge

<p>An edge in a connection.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#person">Person</a>)</td> 
    <td><p>The item at the end of the edge</p></td>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination</p></td>
  </tr>
</table>

---

### StarshipsConnection

<p>A connection to a list of items.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#starshipsedge">[StarshipsEdge]</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>totalCount</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>A count of the total number of objects in this connection, ignoring pagination.
This allows a client to fetch the first five objects by passing &ldquo;5&rdquo; as the
argument to &ldquo;first&rdquo;, then fetch the total count so it could display &ldquo;5 of 83&rdquo;,
for example.</p></td>
  </tr>
  <tr>
    <td><strong>starships</strong> (<a href="objects.md#starship">[Starship]</a>)</td> 
    <td><p>A list of all of the objects returned in the connection. This is a convenience
field provided for quickly exploring the API; rather than querying for
&ldquo;{ edges { node } }&rdquo; when no edge data is needed, this field can be be used
instead. Note that when clients like Relay need to fetch the &ldquo;cursor&rdquo; field on
the edge to enable efficient pagination, this shortcut cannot be used, and the
full &ldquo;{ edges { node } }&rdquo; version should be used instead.</p></td>
  </tr>
</table>

---

### VehiclePilotsEdge

<p>An edge in a connection.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#person">Person</a>)</td> 
    <td><p>The item at the end of the edge</p></td>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination</p></td>
  </tr>
</table>

---

### PersonStarshipsConnection

<p>A connection to a list of items.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#personstarshipsedge">[PersonStarshipsEdge]</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>totalCount</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>A count of the total number of objects in this connection, ignoring pagination.
This allows a client to fetch the first five objects by passing &ldquo;5&rdquo; as the
argument to &ldquo;first&rdquo;, then fetch the total count so it could display &ldquo;5 of 83&rdquo;,
for example.</p></td>
  </tr>
  <tr>
    <td><strong>starships</strong> (<a href="objects.md#starship">[Starship]</a>)</td> 
    <td><p>A list of all of the objects returned in the connection. This is a convenience
field provided for quickly exploring the API; rather than querying for
&ldquo;{ edges { node } }&rdquo; when no edge data is needed, this field can be be used
instead. Note that when clients like Relay need to fetch the &ldquo;cursor&rdquo; field on
the edge to enable efficient pagination, this shortcut cannot be used, and the
full &ldquo;{ edges { node } }&rdquo; version should be used instead.</p></td>
  </tr>
</table>

---

### Planet

<p>A large mass, planet or planetoid in the Star Wars Universe, at the time of
0 ABY.</p> 

#### Implements


- [Node](interfaces.md#node) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>name</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The name of this planet.</p></td>
  </tr>
  <tr>
    <td><strong>diameter</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>The diameter of this planet in kilometers.</p></td>
  </tr>
  <tr>
    <td><strong>rotationPeriod</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>The number of standard hours it takes for this planet to complete a single
rotation on its axis.</p></td>
  </tr>
  <tr>
    <td><strong>orbitalPeriod</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>The number of standard days it takes for this planet to complete a single orbit
of its local star.</p></td>
  </tr>
  <tr>
    <td><strong>gravity</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>A number denoting the gravity of this planet, where &ldquo;1&rdquo; is normal or 1 standard
G. &ldquo;2&rdquo; is twice or 2 standard Gs. &ldquo;0.5&rdquo; is half or 0.5 standard Gs.</p></td>
  </tr>
  <tr>
    <td><strong>population</strong> (<a href="scalars.md#float">Float</a>)</td> 
    <td><p>The average population of sentient beings inhabiting this planet.</p></td>
  </tr>
  <tr>
    <td><strong>climates</strong> (<a href="scalars.md#string">[String]</a>)</td> 
    <td><p>The climates of this planet.</p></td>
  </tr>
  <tr>
    <td><strong>terrains</strong> (<a href="scalars.md#string">[String]</a>)</td> 
    <td><p>The terrains of this planet.</p></td>
  </tr>
  <tr>
    <td><strong>surfaceWater</strong> (<a href="scalars.md#float">Float</a>)</td> 
    <td><p>The percentage of the planet surface that is naturally occuring water or bodies
of water.</p></td>
  </tr>
  <tr>
    <td><strong>residentConnection</strong> (<a href="objects.md#planetresidentsconnection">PlanetResidentsConnection</a>)</td> 
    <td>
      <p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>filmConnection</strong> (<a href="objects.md#planetfilmsconnection">PlanetFilmsConnection</a>)</td> 
    <td>
      <p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>created</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The ISO 8601 date format of the time that this resource was created.</p></td>
  </tr>
  <tr>
    <td><strong>edited</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The ISO 8601 date format of the time that this resource was edited.</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td><p>The ID of an object</p></td>
  </tr>
</table>

---

### SpeciesPeopleConnection

<p>A connection to a list of items.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#speciespeopleedge">[SpeciesPeopleEdge]</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>totalCount</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>A count of the total number of objects in this connection, ignoring pagination.
This allows a client to fetch the first five objects by passing &ldquo;5&rdquo; as the
argument to &ldquo;first&rdquo;, then fetch the total count so it could display &ldquo;5 of 83&rdquo;,
for example.</p></td>
  </tr>
  <tr>
    <td><strong>people</strong> (<a href="objects.md#person">[Person]</a>)</td> 
    <td><p>A list of all of the objects returned in the connection. This is a convenience
field provided for quickly exploring the API; rather than querying for
&ldquo;{ edges { node } }&rdquo; when no edge data is needed, this field can be be used
instead. Note that when clients like Relay need to fetch the &ldquo;cursor&rdquo; field on
the edge to enable efficient pagination, this shortcut cannot be used, and the
full &ldquo;{ edges { node } }&rdquo; version should be used instead.</p></td>
  </tr>
</table>

---

### FilmPlanetsConnection

<p>A connection to a list of items.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#filmplanetsedge">[FilmPlanetsEdge]</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>totalCount</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>A count of the total number of objects in this connection, ignoring pagination.
This allows a client to fetch the first five objects by passing &ldquo;5&rdquo; as the
argument to &ldquo;first&rdquo;, then fetch the total count so it could display &ldquo;5 of 83&rdquo;,
for example.</p></td>
  </tr>
  <tr>
    <td><strong>planets</strong> (<a href="objects.md#planet">[Planet]</a>)</td> 
    <td><p>A list of all of the objects returned in the connection. This is a convenience
field provided for quickly exploring the API; rather than querying for
&ldquo;{ edges { node } }&rdquo; when no edge data is needed, this field can be be used
instead. Note that when clients like Relay need to fetch the &ldquo;cursor&rdquo; field on
the edge to enable efficient pagination, this shortcut cannot be used, and the
full &ldquo;{ edges { node } }&rdquo; version should be used instead.</p></td>
  </tr>
</table>

---

### PlanetsEdge

<p>An edge in a connection.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#planet">Planet</a>)</td> 
    <td><p>The item at the end of the edge</p></td>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination</p></td>
  </tr>
</table>

---

### VehiclePilotsConnection

<p>A connection to a list of items.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#vehiclepilotsedge">[VehiclePilotsEdge]</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>totalCount</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>A count of the total number of objects in this connection, ignoring pagination.
This allows a client to fetch the first five objects by passing &ldquo;5&rdquo; as the
argument to &ldquo;first&rdquo;, then fetch the total count so it could display &ldquo;5 of 83&rdquo;,
for example.</p></td>
  </tr>
  <tr>
    <td><strong>pilots</strong> (<a href="objects.md#person">[Person]</a>)</td> 
    <td><p>A list of all of the objects returned in the connection. This is a convenience
field provided for quickly exploring the API; rather than querying for
&ldquo;{ edges { node } }&rdquo; when no edge data is needed, this field can be be used
instead. Note that when clients like Relay need to fetch the &ldquo;cursor&rdquo; field on
the edge to enable efficient pagination, this shortcut cannot be used, and the
full &ldquo;{ edges { node } }&rdquo; version should be used instead.</p></td>
  </tr>
</table>

---

### FilmPlanetsEdge

<p>An edge in a connection.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#planet">Planet</a>)</td> 
    <td><p>The item at the end of the edge</p></td>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination</p></td>
  </tr>
</table>

---

### PeopleConnection

<p>A connection to a list of items.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#peopleedge">[PeopleEdge]</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>totalCount</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>A count of the total number of objects in this connection, ignoring pagination.
This allows a client to fetch the first five objects by passing &ldquo;5&rdquo; as the
argument to &ldquo;first&rdquo;, then fetch the total count so it could display &ldquo;5 of 83&rdquo;,
for example.</p></td>
  </tr>
  <tr>
    <td><strong>people</strong> (<a href="objects.md#person">[Person]</a>)</td> 
    <td><p>A list of all of the objects returned in the connection. This is a convenience
field provided for quickly exploring the API; rather than querying for
&ldquo;{ edges { node } }&rdquo; when no edge data is needed, this field can be be used
instead. Note that when clients like Relay need to fetch the &ldquo;cursor&rdquo; field on
the edge to enable efficient pagination, this shortcut cannot be used, and the
full &ldquo;{ edges { node } }&rdquo; version should be used instead.</p></td>
  </tr>
</table>

---

### PersonFilmsEdge

<p>An edge in a connection.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#film">Film</a>)</td> 
    <td><p>The item at the end of the edge</p></td>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination</p></td>
  </tr>
</table>

---

### FilmsConnection

<p>A connection to a list of items.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#filmsedge">[FilmsEdge]</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>totalCount</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>A count of the total number of objects in this connection, ignoring pagination.
This allows a client to fetch the first five objects by passing &ldquo;5&rdquo; as the
argument to &ldquo;first&rdquo;, then fetch the total count so it could display &ldquo;5 of 83&rdquo;,
for example.</p></td>
  </tr>
  <tr>
    <td><strong>films</strong> (<a href="objects.md#film">[Film]</a>)</td> 
    <td><p>A list of all of the objects returned in the connection. This is a convenience
field provided for quickly exploring the API; rather than querying for
&ldquo;{ edges { node } }&rdquo; when no edge data is needed, this field can be be used
instead. Note that when clients like Relay need to fetch the &ldquo;cursor&rdquo; field on
the edge to enable efficient pagination, this shortcut cannot be used, and the
full &ldquo;{ edges { node } }&rdquo; version should be used instead.</p></td>
  </tr>
</table>

---

### Starship

<p>A single transport craft that has hyperdrive capability.</p> 

#### Implements


- [Node](interfaces.md#node) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>name</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The name of this starship. The common name, such as &ldquo;Death Star&rdquo;.</p></td>
  </tr>
  <tr>
    <td><strong>model</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The model or official name of this starship. Such as &ldquo;T-65 X-wing&rdquo; or &ldquo;DS-1
Orbital Battle Station&rdquo;.</p></td>
  </tr>
  <tr>
    <td><strong>starshipClass</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The class of this starship, such as &ldquo;Starfighter&rdquo; or &ldquo;Deep Space Mobile
Battlestation&rdquo;</p></td>
  </tr>
  <tr>
    <td><strong>manufacturers</strong> (<a href="scalars.md#string">[String]</a>)</td> 
    <td><p>The manufacturers of this starship.</p></td>
  </tr>
  <tr>
    <td><strong>costInCredits</strong> (<a href="scalars.md#float">Float</a>)</td> 
    <td><p>The cost of this starship new, in galactic credits.</p></td>
  </tr>
  <tr>
    <td><strong>length</strong> (<a href="scalars.md#float">Float</a>)</td> 
    <td><p>The length of this starship in meters.</p></td>
  </tr>
  <tr>
    <td><strong>crew</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The number of personnel needed to run or pilot this starship.</p></td>
  </tr>
  <tr>
    <td><strong>passengers</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The number of non-essential people this starship can transport.</p></td>
  </tr>
  <tr>
    <td><strong>maxAtmospheringSpeed</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>The maximum speed of this starship in atmosphere. null if this starship is
incapable of atmosphering flight.</p></td>
  </tr>
  <tr>
    <td><strong>hyperdriveRating</strong> (<a href="scalars.md#float">Float</a>)</td> 
    <td><p>The class of this starships hyperdrive.</p></td>
  </tr>
  <tr>
    <td><strong>MGLT</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>The Maximum number of Megalights this starship can travel in a standard hour.
A &ldquo;Megalight&rdquo; is a standard unit of distance and has never been defined before
within the Star Wars universe. This figure is only really useful for measuring
the difference in speed of starships. We can assume it is similar to AU, the
distance between our Sun (Sol) and Earth.</p></td>
  </tr>
  <tr>
    <td><strong>cargoCapacity</strong> (<a href="scalars.md#float">Float</a>)</td> 
    <td><p>The maximum number of kilograms that this starship can transport.</p></td>
  </tr>
  <tr>
    <td><strong>consumables</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The maximum length of time that this starship can provide consumables for its
entire crew without having to resupply.</p></td>
  </tr>
  <tr>
    <td><strong>pilotConnection</strong> (<a href="objects.md#starshippilotsconnection">StarshipPilotsConnection</a>)</td> 
    <td>
      <p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>filmConnection</strong> (<a href="objects.md#starshipfilmsconnection">StarshipFilmsConnection</a>)</td> 
    <td>
      <p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>created</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The ISO 8601 date format of the time that this resource was created.</p></td>
  </tr>
  <tr>
    <td><strong>edited</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The ISO 8601 date format of the time that this resource was edited.</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td><p>The ID of an object</p></td>
  </tr>
</table>

---

### VehicleFilmsConnection

<p>A connection to a list of items.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#vehiclefilmsedge">[VehicleFilmsEdge]</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>totalCount</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>A count of the total number of objects in this connection, ignoring pagination.
This allows a client to fetch the first five objects by passing &ldquo;5&rdquo; as the
argument to &ldquo;first&rdquo;, then fetch the total count so it could display &ldquo;5 of 83&rdquo;,
for example.</p></td>
  </tr>
  <tr>
    <td><strong>films</strong> (<a href="objects.md#film">[Film]</a>)</td> 
    <td><p>A list of all of the objects returned in the connection. This is a convenience
field provided for quickly exploring the API; rather than querying for
&ldquo;{ edges { node } }&rdquo; when no edge data is needed, this field can be be used
instead. Note that when clients like Relay need to fetch the &ldquo;cursor&rdquo; field on
the edge to enable efficient pagination, this shortcut cannot be used, and the
full &ldquo;{ edges { node } }&rdquo; version should be used instead.</p></td>
  </tr>
</table>

---

### VehicleFilmsEdge

<p>An edge in a connection.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#film">Film</a>)</td> 
    <td><p>The item at the end of the edge</p></td>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination</p></td>
  </tr>
</table>

---

### VehiclesEdge

<p>An edge in a connection.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#vehicle">Vehicle</a>)</td> 
    <td><p>The item at the end of the edge</p></td>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination</p></td>
  </tr>
</table>

---

### Film

<p>A single film.</p> 

#### Implements


- [Node](interfaces.md#node) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>title</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The title of this film.</p></td>
  </tr>
  <tr>
    <td><strong>episodeID</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>The episode number of this film.</p></td>
  </tr>
  <tr>
    <td><strong>openingCrawl</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The opening paragraphs at the beginning of this film.</p></td>
  </tr>
  <tr>
    <td><strong>director</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The name of the director of this film.</p></td>
  </tr>
  <tr>
    <td><strong>producers</strong> (<a href="scalars.md#string">[String]</a>)</td> 
    <td><p>The name(s) of the producer(s) of this film.</p></td>
  </tr>
  <tr>
    <td><strong>releaseDate</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The ISO 8601 date format of film release at original creator country.</p></td>
  </tr>
  <tr>
    <td><strong>speciesConnection</strong> (<a href="objects.md#filmspeciesconnection">FilmSpeciesConnection</a>)</td> 
    <td>
      <p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>starshipConnection</strong> (<a href="objects.md#filmstarshipsconnection">FilmStarshipsConnection</a>)</td> 
    <td>
      <p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>vehicleConnection</strong> (<a href="objects.md#filmvehiclesconnection">FilmVehiclesConnection</a>)</td> 
    <td>
      <p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>characterConnection</strong> (<a href="objects.md#filmcharactersconnection">FilmCharactersConnection</a>)</td> 
    <td>
      <p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>planetConnection</strong> (<a href="objects.md#filmplanetsconnection">FilmPlanetsConnection</a>)</td> 
    <td>
      <p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>created</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The ISO 8601 date format of the time that this resource was created.</p></td>
  </tr>
  <tr>
    <td><strong>edited</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The ISO 8601 date format of the time that this resource was edited.</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td><p>The ID of an object</p></td>
  </tr>
</table>

---

### FilmStarshipsEdge

<p>An edge in a connection.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#starship">Starship</a>)</td> 
    <td><p>The item at the end of the edge</p></td>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination</p></td>
  </tr>
</table>

---

### StarshipsEdge

<p>An edge in a connection.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#starship">Starship</a>)</td> 
    <td><p>The item at the end of the edge</p></td>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination</p></td>
  </tr>
</table>

---

### FilmStarshipsConnection

<p>A connection to a list of items.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#filmstarshipsedge">[FilmStarshipsEdge]</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>totalCount</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>A count of the total number of objects in this connection, ignoring pagination.
This allows a client to fetch the first five objects by passing &ldquo;5&rdquo; as the
argument to &ldquo;first&rdquo;, then fetch the total count so it could display &ldquo;5 of 83&rdquo;,
for example.</p></td>
  </tr>
  <tr>
    <td><strong>starships</strong> (<a href="objects.md#starship">[Starship]</a>)</td> 
    <td><p>A list of all of the objects returned in the connection. This is a convenience
field provided for quickly exploring the API; rather than querying for
&ldquo;{ edges { node } }&rdquo; when no edge data is needed, this field can be be used
instead. Note that when clients like Relay need to fetch the &ldquo;cursor&rdquo; field on
the edge to enable efficient pagination, this shortcut cannot be used, and the
full &ldquo;{ edges { node } }&rdquo; version should be used instead.</p></td>
  </tr>
</table>

---

### PlanetFilmsEdge

<p>An edge in a connection.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#film">Film</a>)</td> 
    <td><p>The item at the end of the edge</p></td>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination</p></td>
  </tr>
</table>

---

### SpeciesFilmsEdge

<p>An edge in a connection.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#film">Film</a>)</td> 
    <td><p>The item at the end of the edge</p></td>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination</p></td>
  </tr>
</table>

---

### StarshipFilmsEdge

<p>An edge in a connection.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#film">Film</a>)</td> 
    <td><p>The item at the end of the edge</p></td>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination</p></td>
  </tr>
</table>

---

### FilmCharactersEdge

<p>An edge in a connection.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#person">Person</a>)</td> 
    <td><p>The item at the end of the edge</p></td>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination</p></td>
  </tr>
</table>

---

### Person

<p>An individual person or character within the Star Wars universe.</p> 

#### Implements


- [Node](interfaces.md#node) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>name</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The name of this person.</p></td>
  </tr>
  <tr>
    <td><strong>birthYear</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The birth year of the person, using the in-universe standard of BBY or ABY -
Before the Battle of Yavin or After the Battle of Yavin. The Battle of Yavin is
a battle that occurs at the end of Star Wars episode IV: A New Hope.</p></td>
  </tr>
  <tr>
    <td><strong>eyeColor</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The eye color of this person. Will be &ldquo;unknown&rdquo; if not known or &ldquo;n/a&rdquo; if the
person does not have an eye.</p></td>
  </tr>
  <tr>
    <td><strong>gender</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The gender of this person. Either &ldquo;Male&rdquo;, &ldquo;Female&rdquo; or &ldquo;unknown&rdquo;,
&ldquo;n/a&rdquo; if the person does not have a gender.</p></td>
  </tr>
  <tr>
    <td><strong>hairColor</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The hair color of this person. Will be &ldquo;unknown&rdquo; if not known or &ldquo;n/a&rdquo; if the
person does not have hair.</p></td>
  </tr>
  <tr>
    <td><strong>height</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>The height of the person in centimeters.</p></td>
  </tr>
  <tr>
    <td><strong>mass</strong> (<a href="scalars.md#float">Float</a>)</td> 
    <td><p>The mass of the person in kilograms.</p></td>
  </tr>
  <tr>
    <td><strong>skinColor</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The skin color of this person.</p></td>
  </tr>
  <tr>
    <td><strong>homeworld</strong> (<a href="objects.md#planet">Planet</a>)</td> 
    <td><p>A planet that this person was born on or inhabits.</p></td>
  </tr>
  <tr>
    <td><strong>filmConnection</strong> (<a href="objects.md#personfilmsconnection">PersonFilmsConnection</a>)</td> 
    <td>
      <p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>species</strong> (<a href="objects.md#species">Species</a>)</td> 
    <td><p>The species that this person belongs to, or null if unknown.</p></td>
  </tr>
  <tr>
    <td><strong>starshipConnection</strong> (<a href="objects.md#personstarshipsconnection">PersonStarshipsConnection</a>)</td> 
    <td>
      <p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>vehicleConnection</strong> (<a href="objects.md#personvehiclesconnection">PersonVehiclesConnection</a>)</td> 
    <td>
      <p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>created</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The ISO 8601 date format of the time that this resource was created.</p></td>
  </tr>
  <tr>
    <td><strong>edited</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The ISO 8601 date format of the time that this resource was edited.</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td><p>The ID of an object</p></td>
  </tr>
</table>

---

### PlanetResidentsEdge

<p>An edge in a connection.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#person">Person</a>)</td> 
    <td><p>The item at the end of the edge</p></td>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination</p></td>
  </tr>
</table>

---

### Species

<p>A type of person or character within the Star Wars Universe.</p> 

#### Implements


- [Node](interfaces.md#node) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>name</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The name of this species.</p></td>
  </tr>
  <tr>
    <td><strong>classification</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The classification of this species, such as &ldquo;mammal&rdquo; or &ldquo;reptile&rdquo;.</p></td>
  </tr>
  <tr>
    <td><strong>designation</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The designation of this species, such as &ldquo;sentient&rdquo;.</p></td>
  </tr>
  <tr>
    <td><strong>averageHeight</strong> (<a href="scalars.md#float">Float</a>)</td> 
    <td><p>The average height of this species in centimeters.</p></td>
  </tr>
  <tr>
    <td><strong>averageLifespan</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>The average lifespan of this species in years, null if unknown.</p></td>
  </tr>
  <tr>
    <td><strong>eyeColors</strong> (<a href="scalars.md#string">[String]</a>)</td> 
    <td><p>Common eye colors for this species, null if this species does not typically
have eyes.</p></td>
  </tr>
  <tr>
    <td><strong>hairColors</strong> (<a href="scalars.md#string">[String]</a>)</td> 
    <td><p>Common hair colors for this species, null if this species does not typically
have hair.</p></td>
  </tr>
  <tr>
    <td><strong>skinColors</strong> (<a href="scalars.md#string">[String]</a>)</td> 
    <td><p>Common skin colors for this species, null if this species does not typically
have skin.</p></td>
  </tr>
  <tr>
    <td><strong>language</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The language commonly spoken by this species.</p></td>
  </tr>
  <tr>
    <td><strong>homeworld</strong> (<a href="objects.md#planet">Planet</a>)</td> 
    <td><p>A planet that this species originates from.</p></td>
  </tr>
  <tr>
    <td><strong>personConnection</strong> (<a href="objects.md#speciespeopleconnection">SpeciesPeopleConnection</a>)</td> 
    <td>
      <p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>filmConnection</strong> (<a href="objects.md#speciesfilmsconnection">SpeciesFilmsConnection</a>)</td> 
    <td>
      <p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>created</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The ISO 8601 date format of the time that this resource was created.</p></td>
  </tr>
  <tr>
    <td><strong>edited</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The ISO 8601 date format of the time that this resource was edited.</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td><p>The ID of an object</p></td>
  </tr>
</table>

---

### SpeciesConnection

<p>A connection to a list of items.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#speciesedge">[SpeciesEdge]</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>totalCount</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>A count of the total number of objects in this connection, ignoring pagination.
This allows a client to fetch the first five objects by passing &ldquo;5&rdquo; as the
argument to &ldquo;first&rdquo;, then fetch the total count so it could display &ldquo;5 of 83&rdquo;,
for example.</p></td>
  </tr>
  <tr>
    <td><strong>species</strong> (<a href="objects.md#species">[Species]</a>)</td> 
    <td><p>A list of all of the objects returned in the connection. This is a convenience
field provided for quickly exploring the API; rather than querying for
&ldquo;{ edges { node } }&rdquo; when no edge data is needed, this field can be be used
instead. Note that when clients like Relay need to fetch the &ldquo;cursor&rdquo; field on
the edge to enable efficient pagination, this shortcut cannot be used, and the
full &ldquo;{ edges { node } }&rdquo; version should be used instead.</p></td>
  </tr>
</table>

---

### PersonStarshipsEdge

<p>An edge in a connection.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#starship">Starship</a>)</td> 
    <td><p>The item at the end of the edge</p></td>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination</p></td>
  </tr>
</table>

---

### Root

  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>allFilms</strong> (<a href="objects.md#filmsconnection">FilmsConnection</a>)</td> 
    <td>
      <p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>film</strong> (<a href="objects.md#film">Film</a>)</td> 
    <td>
      <p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>id (<a href="scalars.md#id">ID</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>filmID (<a href="scalars.md#id">ID</a>)</p>
            <p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>allPeople</strong> (<a href="objects.md#peopleconnection">PeopleConnection</a>)</td> 
    <td>
      <p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>person</strong> (<a href="objects.md#person">Person</a>)</td> 
    <td>
      <p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>id (<a href="scalars.md#id">ID</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>personID (<a href="scalars.md#id">ID</a>)</p>
            <p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>allPlanets</strong> (<a href="objects.md#planetsconnection">PlanetsConnection</a>)</td> 
    <td>
      <p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>planet</strong> (<a href="objects.md#planet">Planet</a>)</td> 
    <td>
      <p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>id (<a href="scalars.md#id">ID</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>planetID (<a href="scalars.md#id">ID</a>)</p>
            <p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>allSpecies</strong> (<a href="objects.md#speciesconnection">SpeciesConnection</a>)</td> 
    <td>
      <p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>species</strong> (<a href="objects.md#species">Species</a>)</td> 
    <td>
      <p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>id (<a href="scalars.md#id">ID</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>speciesID (<a href="scalars.md#id">ID</a>)</p>
            <p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>allStarships</strong> (<a href="objects.md#starshipsconnection">StarshipsConnection</a>)</td> 
    <td>
      <p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>starship</strong> (<a href="objects.md#starship">Starship</a>)</td> 
    <td>
      <p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>id (<a href="scalars.md#id">ID</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>starshipID (<a href="scalars.md#id">ID</a>)</p>
            <p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>allVehicles</strong> (<a href="objects.md#vehiclesconnection">VehiclesConnection</a>)</td> 
    <td>
      <p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>vehicle</strong> (<a href="objects.md#vehicle">Vehicle</a>)</td> 
    <td>
      <p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>id (<a href="scalars.md#id">ID</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>vehicleID (<a href="scalars.md#id">ID</a>)</p>
            <p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="interfaces.md#node">Node</a>)</td> 
    <td>
      <p><p>Fetches an object given its ID</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>id (<a href="scalars.md#id">ID!</a>)</p>
            <p><p>The ID of an object</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
</table>

---

### SpeciesEdge

<p>An edge in a connection.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#species">Species</a>)</td> 
    <td><p>The item at the end of the edge</p></td>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination</p></td>
  </tr>
</table>

---

### PersonVehiclesConnection

<p>A connection to a list of items.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#personvehiclesedge">[PersonVehiclesEdge]</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>totalCount</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>A count of the total number of objects in this connection, ignoring pagination.
This allows a client to fetch the first five objects by passing &ldquo;5&rdquo; as the
argument to &ldquo;first&rdquo;, then fetch the total count so it could display &ldquo;5 of 83&rdquo;,
for example.</p></td>
  </tr>
  <tr>
    <td><strong>vehicles</strong> (<a href="objects.md#vehicle">[Vehicle]</a>)</td> 
    <td><p>A list of all of the objects returned in the connection. This is a convenience
field provided for quickly exploring the API; rather than querying for
&ldquo;{ edges { node } }&rdquo; when no edge data is needed, this field can be be used
instead. Note that when clients like Relay need to fetch the &ldquo;cursor&rdquo; field on
the edge to enable efficient pagination, this shortcut cannot be used, and the
full &ldquo;{ edges { node } }&rdquo; version should be used instead.</p></td>
  </tr>
</table>

---

### PersonVehiclesEdge

<p>An edge in a connection.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#vehicle">Vehicle</a>)</td> 
    <td><p>The item at the end of the edge</p></td>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination</p></td>
  </tr>
</table>

---

### FilmVehiclesEdge

<p>An edge in a connection.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#vehicle">Vehicle</a>)</td> 
    <td><p>The item at the end of the edge</p></td>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination</p></td>
  </tr>
</table>

---

### StarshipPilotsEdge

<p>An edge in a connection.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#person">Person</a>)</td> 
    <td><p>The item at the end of the edge</p></td>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination</p></td>
  </tr>
</table>

---

### StarshipFilmsConnection

<p>A connection to a list of items.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#starshipfilmsedge">[StarshipFilmsEdge]</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>totalCount</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>A count of the total number of objects in this connection, ignoring pagination.
This allows a client to fetch the first five objects by passing &ldquo;5&rdquo; as the
argument to &ldquo;first&rdquo;, then fetch the total count so it could display &ldquo;5 of 83&rdquo;,
for example.</p></td>
  </tr>
  <tr>
    <td><strong>films</strong> (<a href="objects.md#film">[Film]</a>)</td> 
    <td><p>A list of all of the objects returned in the connection. This is a convenience
field provided for quickly exploring the API; rather than querying for
&ldquo;{ edges { node } }&rdquo; when no edge data is needed, this field can be be used
instead. Note that when clients like Relay need to fetch the &ldquo;cursor&rdquo; field on
the edge to enable efficient pagination, this shortcut cannot be used, and the
full &ldquo;{ edges { node } }&rdquo; version should be used instead.</p></td>
  </tr>
</table>

---

### Vehicle

<p>A single transport craft that does not have hyperdrive capability</p> 

#### Implements


- [Node](interfaces.md#node) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>name</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The name of this vehicle. The common name, such as &ldquo;Sand Crawler&rdquo; or &ldquo;Speeder
bike&rdquo;.</p></td>
  </tr>
  <tr>
    <td><strong>model</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The model or official name of this vehicle. Such as &ldquo;All-Terrain Attack
Transport&rdquo;.</p></td>
  </tr>
  <tr>
    <td><strong>vehicleClass</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The class of this vehicle, such as &ldquo;Wheeled&rdquo; or &ldquo;Repulsorcraft&rdquo;.</p></td>
  </tr>
  <tr>
    <td><strong>manufacturers</strong> (<a href="scalars.md#string">[String]</a>)</td> 
    <td><p>The manufacturers of this vehicle.</p></td>
  </tr>
  <tr>
    <td><strong>costInCredits</strong> (<a href="scalars.md#float">Float</a>)</td> 
    <td><p>The cost of this vehicle new, in Galactic Credits.</p></td>
  </tr>
  <tr>
    <td><strong>length</strong> (<a href="scalars.md#float">Float</a>)</td> 
    <td><p>The length of this vehicle in meters.</p></td>
  </tr>
  <tr>
    <td><strong>crew</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The number of personnel needed to run or pilot this vehicle.</p></td>
  </tr>
  <tr>
    <td><strong>passengers</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The number of non-essential people this vehicle can transport.</p></td>
  </tr>
  <tr>
    <td><strong>maxAtmospheringSpeed</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>The maximum speed of this vehicle in atmosphere.</p></td>
  </tr>
  <tr>
    <td><strong>cargoCapacity</strong> (<a href="scalars.md#float">Float</a>)</td> 
    <td><p>The maximum number of kilograms that this vehicle can transport.</p></td>
  </tr>
  <tr>
    <td><strong>consumables</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The maximum length of time that this vehicle can provide consumables for its
entire crew without having to resupply.</p></td>
  </tr>
  <tr>
    <td><strong>pilotConnection</strong> (<a href="objects.md#vehiclepilotsconnection">VehiclePilotsConnection</a>)</td> 
    <td>
      <p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>filmConnection</strong> (<a href="objects.md#vehiclefilmsconnection">VehicleFilmsConnection</a>)</td> 
    <td>
      <p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>created</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The ISO 8601 date format of the time that this resource was created.</p></td>
  </tr>
  <tr>
    <td><strong>edited</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The ISO 8601 date format of the time that this resource was edited.</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td><p>The ID of an object</p></td>
  </tr>
</table>

---

### FilmSpeciesConnection

<p>A connection to a list of items.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#filmspeciesedge">[FilmSpeciesEdge]</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>totalCount</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>A count of the total number of objects in this connection, ignoring pagination.
This allows a client to fetch the first five objects by passing &ldquo;5&rdquo; as the
argument to &ldquo;first&rdquo;, then fetch the total count so it could display &ldquo;5 of 83&rdquo;,
for example.</p></td>
  </tr>
  <tr>
    <td><strong>species</strong> (<a href="objects.md#species">[Species]</a>)</td> 
    <td><p>A list of all of the objects returned in the connection. This is a convenience
field provided for quickly exploring the API; rather than querying for
&ldquo;{ edges { node } }&rdquo; when no edge data is needed, this field can be be used
instead. Note that when clients like Relay need to fetch the &ldquo;cursor&rdquo; field on
the edge to enable efficient pagination, this shortcut cannot be used, and the
full &ldquo;{ edges { node } }&rdquo; version should be used instead.</p></td>
  </tr>
</table>

---

### FilmSpeciesEdge

<p>An edge in a connection.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#species">Species</a>)</td> 
    <td><p>The item at the end of the edge</p></td>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination</p></td>
  </tr>
</table>

---

### FilmVehiclesConnection

<p>A connection to a list of items.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#filmvehiclesedge">[FilmVehiclesEdge]</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>totalCount</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>A count of the total number of objects in this connection, ignoring pagination.
This allows a client to fetch the first five objects by passing &ldquo;5&rdquo; as the
argument to &ldquo;first&rdquo;, then fetch the total count so it could display &ldquo;5 of 83&rdquo;,
for example.</p></td>
  </tr>
  <tr>
    <td><strong>vehicles</strong> (<a href="objects.md#vehicle">[Vehicle]</a>)</td> 
    <td><p>A list of all of the objects returned in the connection. This is a convenience
field provided for quickly exploring the API; rather than querying for
&ldquo;{ edges { node } }&rdquo; when no edge data is needed, this field can be be used
instead. Note that when clients like Relay need to fetch the &ldquo;cursor&rdquo; field on
the edge to enable efficient pagination, this shortcut cannot be used, and the
full &ldquo;{ edges { node } }&rdquo; version should be used instead.</p></td>
  </tr>
</table>

---

### PeopleEdge

<p>An edge in a connection.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#person">Person</a>)</td> 
    <td><p>The item at the end of the edge</p></td>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination</p></td>
  </tr>
</table>

---

### PlanetResidentsConnection

<p>A connection to a list of items.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#planetresidentsedge">[PlanetResidentsEdge]</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>totalCount</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>A count of the total number of objects in this connection, ignoring pagination.
This allows a client to fetch the first five objects by passing &ldquo;5&rdquo; as the
argument to &ldquo;first&rdquo;, then fetch the total count so it could display &ldquo;5 of 83&rdquo;,
for example.</p></td>
  </tr>
  <tr>
    <td><strong>residents</strong> (<a href="objects.md#person">[Person]</a>)</td> 
    <td><p>A list of all of the objects returned in the connection. This is a convenience
field provided for quickly exploring the API; rather than querying for
&ldquo;{ edges { node } }&rdquo; when no edge data is needed, this field can be be used
instead. Note that when clients like Relay need to fetch the &ldquo;cursor&rdquo; field on
the edge to enable efficient pagination, this shortcut cannot be used, and the
full &ldquo;{ edges { node } }&rdquo; version should be used instead.</p></td>
  </tr>
</table>

---

### PlanetsConnection

<p>A connection to a list of items.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#planetsedge">[PlanetsEdge]</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>totalCount</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>A count of the total number of objects in this connection, ignoring pagination.
This allows a client to fetch the first five objects by passing &ldquo;5&rdquo; as the
argument to &ldquo;first&rdquo;, then fetch the total count so it could display &ldquo;5 of 83&rdquo;,
for example.</p></td>
  </tr>
  <tr>
    <td><strong>planets</strong> (<a href="objects.md#planet">[Planet]</a>)</td> 
    <td><p>A list of all of the objects returned in the connection. This is a convenience
field provided for quickly exploring the API; rather than querying for
&ldquo;{ edges { node } }&rdquo; when no edge data is needed, this field can be be used
instead. Note that when clients like Relay need to fetch the &ldquo;cursor&rdquo; field on
the edge to enable efficient pagination, this shortcut cannot be used, and the
full &ldquo;{ edges { node } }&rdquo; version should be used instead.</p></td>
  </tr>
</table>

---

### VehiclesConnection

<p>A connection to a list of items.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#vehiclesedge">[VehiclesEdge]</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>totalCount</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>A count of the total number of objects in this connection, ignoring pagination.
This allows a client to fetch the first five objects by passing &ldquo;5&rdquo; as the
argument to &ldquo;first&rdquo;, then fetch the total count so it could display &ldquo;5 of 83&rdquo;,
for example.</p></td>
  </tr>
  <tr>
    <td><strong>vehicles</strong> (<a href="objects.md#vehicle">[Vehicle]</a>)</td> 
    <td><p>A list of all of the objects returned in the connection. This is a convenience
field provided for quickly exploring the API; rather than querying for
&ldquo;{ edges { node } }&rdquo; when no edge data is needed, this field can be be used
instead. Note that when clients like Relay need to fetch the &ldquo;cursor&rdquo; field on
the edge to enable efficient pagination, this shortcut cannot be used, and the
full &ldquo;{ edges { node } }&rdquo; version should be used instead.</p></td>
  </tr>
</table>

---

### FilmCharactersConnection

<p>A connection to a list of items.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#filmcharactersedge">[FilmCharactersEdge]</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>totalCount</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>A count of the total number of objects in this connection, ignoring pagination.
This allows a client to fetch the first five objects by passing &ldquo;5&rdquo; as the
argument to &ldquo;first&rdquo;, then fetch the total count so it could display &ldquo;5 of 83&rdquo;,
for example.</p></td>
  </tr>
  <tr>
    <td><strong>characters</strong> (<a href="objects.md#person">[Person]</a>)</td> 
    <td><p>A list of all of the objects returned in the connection. This is a convenience
field provided for quickly exploring the API; rather than querying for
&ldquo;{ edges { node } }&rdquo; when no edge data is needed, this field can be be used
instead. Note that when clients like Relay need to fetch the &ldquo;cursor&rdquo; field on
the edge to enable efficient pagination, this shortcut cannot be used, and the
full &ldquo;{ edges { node } }&rdquo; version should be used instead.</p></td>
  </tr>
</table>

---

### FilmsEdge

<p>An edge in a connection.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#film">Film</a>)</td> 
    <td><p>The item at the end of the edge</p></td>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination</p></td>
  </tr>
</table>

---

### PlanetFilmsConnection

<p>A connection to a list of items.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#planetfilmsedge">[PlanetFilmsEdge]</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>totalCount</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>A count of the total number of objects in this connection, ignoring pagination.
This allows a client to fetch the first five objects by passing &ldquo;5&rdquo; as the
argument to &ldquo;first&rdquo;, then fetch the total count so it could display &ldquo;5 of 83&rdquo;,
for example.</p></td>
  </tr>
  <tr>
    <td><strong>films</strong> (<a href="objects.md#film">[Film]</a>)</td> 
    <td><p>A list of all of the objects returned in the connection. This is a convenience
field provided for quickly exploring the API; rather than querying for
&ldquo;{ edges { node } }&rdquo; when no edge data is needed, this field can be be used
instead. Note that when clients like Relay need to fetch the &ldquo;cursor&rdquo; field on
the edge to enable efficient pagination, this shortcut cannot be used, and the
full &ldquo;{ edges { node } }&rdquo; version should be used instead.</p></td>
  </tr>
</table>

---

### StarshipPilotsConnection

<p>A connection to a list of items.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#starshippilotsedge">[StarshipPilotsEdge]</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>totalCount</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>A count of the total number of objects in this connection, ignoring pagination.
This allows a client to fetch the first five objects by passing &ldquo;5&rdquo; as the
argument to &ldquo;first&rdquo;, then fetch the total count so it could display &ldquo;5 of 83&rdquo;,
for example.</p></td>
  </tr>
  <tr>
    <td><strong>pilots</strong> (<a href="objects.md#person">[Person]</a>)</td> 
    <td><p>A list of all of the objects returned in the connection. This is a convenience
field provided for quickly exploring the API; rather than querying for
&ldquo;{ edges { node } }&rdquo; when no edge data is needed, this field can be be used
instead. Note that when clients like Relay need to fetch the &ldquo;cursor&rdquo; field on
the edge to enable efficient pagination, this shortcut cannot be used, and the
full &ldquo;{ edges { node } }&rdquo; version should be used instead.</p></td>
  </tr>
</table>

---