function Ping() {
  console.log('/ping was hit successfully !')
  var response = `Hi there ! 👋 My host-id is unknown 🕸️`
  return (
    <div>
      <span>
        {response}
      </span>
    </div>
  );
}

export default Ping;
