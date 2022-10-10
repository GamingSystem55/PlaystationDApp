import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { AmdJaguarX86ChipComponent } from './amd-jaguar-x86-chip/amd-jaguar-x86-chip.component';
import { AmdJaguarX86CUComponent } from './amd-jaguar-x86-cu/amd-jaguar-x86-cu.component';
import { AmdJaguarX86ALUComponent } from './amd-jaguar-x86-alu/amd-jaguar-x86-alu.component';

@NgModule({
  declarations: [
    AppComponent,
    AmdJaguarX86ChipComponent,
    AmdJaguarX86CUComponent,
    AmdJaguarX86ALUComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
